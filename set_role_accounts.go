package gconfig

import (
	"fmt"
	"regexp"
	"strings"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
)

type accountAndProvider struct {
	Account  *gconfigv1alpha1.Account
	Provider *gconfigv1alpha1.Provider
}

// this will return true if the format is a string containing 12 numeric characters
// if false, we can treat this account as an alias
func IsStringAnAWSAccountID(input string) bool {
	if len(input) != 12 {
		return false
	}
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	return strings.IndexFunc(string(input), isNotDigit) == -1
}

func IsStringAnAWSOUID(input string) bool {
	// "ou-4w0n-bads234"

	if len(input) != 15 {
		return false
	}
	// This only returns an error if the regex doesn't compile, so we ignore it
	matched, _ := regexp.MatchString("ou-([a-z0-9]){4}-([a-z0-9]){7}", input)

	return matched
}

// setRoleAccounts sets the RoleAccounts field on all Roles in the config.
// It should be called as part of parsing config.
//
// We allow Granted users to specify config using just an account ID or an alias,
// rather than specifying both the provider and the account.
// setRoleAccounts looks up the account string that the user has used against
// our providers.
func (c *Config) setRoleAccounts() error {
	accountMap := make(map[string]accountAndProvider)
	aliasMap := make(map[string][]accountAndProvider)
	for _, p := range c.providers.Providers {
		switch v := p.Details.(type) {
		case *gconfigv1alpha1.Provider_Aws:
			for _, acc := range v.Aws.Accounts {
				collectAccountAndProvider(acc, p, accountMap, aliasMap)
			}
		case *gconfigv1alpha1.Provider_AwsSso:
			for _, acc := range v.AwsSso.Accounts {
				collectAccountAndProvider(acc, p, accountMap, aliasMap)
			}
		}
	}

	for _, r := range c.Roles {
		for _, a := range r.Accounts {
			// logic for matching aliases
			accountPieces := strings.Split(a.Account, ":")
			numberOfPieces := len(accountPieces)
			accountId := a.Account
			if numberOfPieces > 3 {
				err := fmt.Errorf("role %s references an account that is in the wrong format: %s . \naccount must be in the format <accountId> or <alias> or <provider>:<alias> or <provider>:<alias>:<accountId>", r.ID, a.Account)
				err = printLintError(r, err)
				return err
			} else if numberOfPieces == 3 {
				// an account in the format <provider>:<alias>:<accountId> we can use the account id directly
				accountId = accountPieces[2]
			} else {
				// if its 2 parts or its not an accountid it must be an alias
				if numberOfPieces == 2 || !(IsStringAnAWSAccountID(a.Account) || IsStringAnAWSOUID(a.Account)) {
					alias := a.Account
					if numberOfPieces == 2 {
						providerExists := false
						for _, p := range c.providers.Providers {
							if p.Id == accountPieces[0] {
								providerExists = true
								break
							}
						}
						if !providerExists {
							err := fmt.Errorf("role %s references a provider that doesn't exist: %s \naccount must be in the format <accountId> or <alias> or <provider>:<alias> or <provider>:<alias>:<accountId>", r.ID, accountPieces[0])
							err = printLintError(r, err)
							return err
						}
						alias = accountPieces[1]
					}
					// it must be an alias
					// find a match then set the acountid =
					aliasAccounts, ok := aliasMap[alias]
					if !ok {
						err := fmt.Errorf("role %s references an account alias that doesn't exist: %s \naccount must be in the format <accountId> or <alias> or <provider>:<alias> or <provider>:<alias>:<accountId>", r.ID, a.Account)
						err = printLintError(r, err)
						return err
					}
					if len(aliasAccounts) > 1 {
						if numberOfPieces == 2 {
							pacc := []accountAndProvider{}
							for _, ac := range aliasAccounts {
								if ac.Provider.Id == accountPieces[0] {
									pacc = append(pacc, ac)
								}
							}
							if len(pacc) > 1 {
								err := generateAmbiguousAliasError(r, a.Account, alias, pacc)
								err = printLintError(r, err)
								return err
							} else if len(pacc) == 1 {
								accountId = pacc[0].Account.Id
							} else {
								err := fmt.Errorf("role %s references an account alias that doesn't exist for this provider: %s \naccount must be in the format <accountId> or <alias> or <provider>:<alias> or <provider>:<alias>:<accountId>", r.ID, a.Account)
								err = printLintError(r, err)
								return err
							}

						} else {
							err := generateAmbiguousAliasError(r, a.Account, alias, aliasAccounts)
							err = printLintError(r, err)
							return err
						}
					} else {
						// only one alias macthes so we use that as the account
						accountId = aliasAccounts[0].Account.Id
					}
				}
			}
			ap, ok := accountMap[accountId]
			if !ok {
				err := fmt.Errorf("role %s references an account that doesn't exist: %s", r.ID, a.Account)
				err = printLintError(r, err)
				return err
			}
			// Configure the default region in order of precedence account > role > provider
			// default region can still end up empty, in which case it is up to the end user to set a region e.g the cli
			defaultRegion := a.DefaultRegion
			if defaultRegion == "" {
				if r.DefaultRegion != "" {
					defaultRegion = r.DefaultRegion
				} else {
					for _, po := range c.ProviderOverrides {
						if po.ID == ap.Provider.Id && po.DefaultRegion != "" {
							defaultRegion = po.DefaultRegion
							break
						}
					}
				}
			}
			if ap.Account.Type == gconfigv1alpha1.Account_TYPE_UNSPECIFIED {
				// if it is an OU account, add all the children rather than the ou
				for _, acc := range ap.Account.Children {
					r.roleAccounts = append(r.roleAccounts, RoleAccount{
						AccountID:     acc.Id,
						ProviderID:    ap.Provider.Id,
						DefaultRegion: defaultRegion,
					})
				}
			} else {
				r.roleAccounts = append(r.roleAccounts, RoleAccount{
					AccountID:     ap.Account.Id,
					ProviderID:    ap.Provider.Id,
					DefaultRegion: defaultRegion,
				})
			}

		}
	}
	return nil
}

func collectAccountAndProvider(a *gconfigv1alpha1.Account, p *gconfigv1alpha1.Provider, accountMap map[string]accountAndProvider, aliasMap map[string][]accountAndProvider) {
	ap := accountAndProvider{
		Account:  a,
		Provider: p,
	}
	accountMap[a.Id] = ap
	if a.Name != "" {
		aliasMap[a.Name] = append(aliasMap[a.Name], ap)
	}
	for _, alias := range a.Aliases {
		if alias != "" {
			aliasMap[alias] = append(aliasMap[alias], ap)
		}
	}

	for _, child := range a.Children {
		collectAccountAndProvider(child, p, accountMap, aliasMap)
	}
}

func generateAmbiguousAliasError(r *Role, a string, alias string, aliasAccounts []accountAndProvider) error {
	errMessage := ""
	if r == nil {
		errMessage = fmt.Sprintf("account '%s' is ambiguous and could refer to one of these account names:\n\n", a)
	} else {
		errMessage = fmt.Sprintf("role %s: account '%s' is ambiguous and could refer to one of these account names:\n\n", r.ID, a)
	}

	example := ""
	for i, account := range aliasAccounts {
		if i == 0 {
			example = fmt.Sprintf("%s:%s:%s", account.Provider.Id, alias, account.Account.Id)
		}
		errMessage += fmt.Sprintf("    - %s:%s:%s (%s %s in provider %s)\n", account.Provider.Id, alias, account.Account.Id, account.Account.Type.String(), account.Account.Id, account.Provider.Id)
	}
	// 	- aws:dev:ou-4w0n-bads234    (AWS Org Unit ou-4w0n-bads234 in provider aws)
	//   - aws:dev:123456789012       (AWS Account 123456789012 in provider aws)

	errMessage += fmt.Sprintf("\nPlease replace '%s' with the account name above that you meant (e.g. %s).", a, example)

	return fmt.Errorf(errMessage)
}

// This utility is used in the CLI to match an alias or account string to a provider and account in the config
func MatchAccountOrAlias(providers []*gconfigv1alpha1.Provider, accountInput string) (*accountAndProvider, error) {
	accountMap := make(map[string]accountAndProvider)
	aliasMap := make(map[string][]accountAndProvider)
	for _, p := range providers {
		switch v := p.Details.(type) {
		case *gconfigv1alpha1.Provider_Aws:
			for _, acc := range v.Aws.Accounts {
				collectAccountAndProvider(acc, p, accountMap, aliasMap)
			}
		case *gconfigv1alpha1.Provider_AwsSso:
			for _, acc := range v.AwsSso.Accounts {
				collectAccountAndProvider(acc, p, accountMap, aliasMap)
			}
		}
	}

	// logic for matching aliases
	accountPieces := strings.Split(accountInput, ":")
	numberOfPieces := len(accountPieces)
	accountId := accountInput
	if numberOfPieces > 3 {
		return nil, fmt.Errorf("account: %s, must be in the format <accountId> or <alias> or <provider>:<alias> or <provider>:<alias>:<accountId>", accountInput)

	} else if numberOfPieces == 3 {
		// an account in the format <provider>:<alias>:<accountId> we can use the account id directly
		accountId = accountPieces[2]
	} else {
		if numberOfPieces == 2 || !(IsStringAnAWSAccountID(accountInput) || IsStringAnAWSOUID(accountInput)) {
			alias := accountInput
			if numberOfPieces == 2 {
				alias = accountPieces[1]
			}
			// it must be an alias
			// find a match then set the acountid =
			aliasAccounts, ok := aliasMap[alias]
			if !ok {
				return nil, fmt.Errorf("account alias does not exist: %s\naccount must be in the format <accountId> or <alias> or <provider>:<alias> or <provider>:<alias>:<accountId>", accountInput)
			}
			if len(aliasAccounts) > 1 {
				return nil, generateAmbiguousAliasError(nil, accountInput, alias, aliasAccounts)

			}
			// only one alias matches so we use that as the account
			accountId = aliasAccounts[0].Account.Id
		}

	}
	ap, ok := accountMap[accountId]
	if !ok {
		return nil, fmt.Errorf("account alias does not exist: %s", accountInput)
	}

	// If an account is an OU, then return the account does not exist error because it cannot be assumed
	if ap.Account.Type == gconfigv1alpha1.Account_TYPE_UNSPECIFIED {
		return nil, fmt.Errorf("account alias does not exist: %s", accountInput)
	} else {
		return &ap, nil
	}
}
