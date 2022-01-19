package gconfig

import (
	"fmt"
	"strings"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
)

type accountAndProvider struct {
	Account  *gconfigv1alpha1.Account
	Provider *gconfigv1alpha1.Provider
}

// this will return true if the format is a string containing 12 numeric characters
// if false, we can treat this account as an alias
func IsAccountAnAWSAccountID(account string) bool {
	if len(account) != 12 {
		return false
	}
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	return strings.IndexFunc(string(account), isNotDigit) == -1
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
		for _, acc := range p.Accounts {
			collectAccountAndProvider(acc, p, accountMap, aliasMap)
		}
	}

	for _, r := range c.Roles {
		for _, a := range r.Accounts {
			// logic for matching aliases
			accountPieces := strings.Split(a, ":")
			numberOfPieces := len(accountPieces)
			accountId := a
			var ap accountAndProvider
			if numberOfPieces > 3 {
				err := fmt.Errorf("role %s references an account that is in the wrong format: %s . Account must be in the format(s) <accountId> or <alias> or <provider>:<alias> or <provider>:<alias>:<accountId>", r.ID, a)
				err = printLintError(r, err)
				return err
			} else if numberOfPieces == 3 {
				accountId = accountPieces[2]
			} else {
				if numberOfPieces == 2 || !IsAccountAnAWSAccountID(a) {
					alias := a
					if numberOfPieces == 2 {
						alias = accountPieces[1]
					}
					// it must be an alias
					// find a match then set the acountid =
					aliasAccounts, ok := aliasMap[alias]
					if !ok {
						err := fmt.Errorf("role %s references an account alias that doesn't exist: %s", r.ID, a)
						err = printLintError(r, err)
						return err
					}
					if len(aliasAccounts) > 1 {
						err := generateAmbiguousAliasError(r, a, alias, aliasAccounts)
						err = printLintError(r, err)
						return err

					}
					// only one alias macthes so we use that as the account
					accountId = aliasAccounts[0].Account.Id
				}

			}
			v, ok := accountMap[accountId]
			if !ok {
				err := fmt.Errorf("role %s references an account that doesn't exist: %s", r.ID, a)
				err = printLintError(r, err)
				return err
			}
			ap = v

			ra := RoleAccount{
				AccountID:  ap.Account.Id,
				ProviderID: ap.Provider.Id,
			}
			r.roleAccounts = append(r.roleAccounts, ra)
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
	errMessage := fmt.Sprintf("role %s: account '%s' is ambiguous and could refer to one of these account names:\n\n", r.ID, a)
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
