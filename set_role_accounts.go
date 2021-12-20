package gconfig

import (
	"fmt"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
)

type accountAndProvider struct {
	Account  *gconfigv1alpha1.Account
	Provider *gconfigv1alpha1.Provider
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

	for _, p := range c.providers.Providers {
		for _, acc := range p.Accounts {
			collectAccountAndProvider(acc, p, accountMap)
		}
	}

	for _, r := range c.Roles {
		for _, a := range r.Accounts {
			v, ok := accountMap[a]
			if !ok {
				err := fmt.Errorf("role %s references an account that doesn't exist: %s", r.ID, a)
				err = printLintError(r, err)
				return err
			}
			ra := RoleAccount{
				AccountID:  v.Account.Id,
				ProviderID: v.Provider.Id,
			}
			r.roleAccounts = append(r.roleAccounts, ra)
		}
	}
	return nil
}

func collectAccountAndProvider(a *gconfigv1alpha1.Account, p *gconfigv1alpha1.Provider, accountMap map[string]accountAndProvider) {
	accountMap[a.Id] = accountAndProvider{
		Account:  a,
		Provider: p,
	}
	for _, child := range a.Children {
		collectAccountAndProvider(child, p, accountMap)
	}
}
