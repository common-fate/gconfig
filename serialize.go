package gconfig

import (
	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"google.golang.org/protobuf/types/known/durationpb"
)

func (c *Config) SerializeProtobuf() *gconfigv1alpha1.Config {
	out := &gconfigv1alpha1.Config{}

	for _, u := range c.Admins {
		out.Admins = append(out.Admins, &gconfigv1alpha1.Member{
			Email: u.Email,
		})
	}
	for _, u := range c.Users {
		out.Users = append(out.Users, &gconfigv1alpha1.Member{
			Email: u.Email,
		})
	}
	for _, g := range c.Groups {
		group := &gconfigv1alpha1.Group{
			Name: g.Name,
			Id:   g.ID,
		}
		for _, u := range g.Members {
			group.Members = append(group.Members, &gconfigv1alpha1.Member{
				Email: u.Email,
			})
		}
		out.Groups = append(out.Groups, group)
	}
	for _, r := range c.Roles {
		role := &gconfigv1alpha1.Role{
			Id:              r.ID,
			Policy:          r.Policy,
			SessionDuration: durationpb.New(r.SessionDuration),
			Group:           r.Group,
			Type:            gconfigv1alpha1.RoleType(gconfigv1alpha1.RoleType_value[r.Type]),
		}
		for _, ra := range r.roleAccounts {
			role.Accounts = append(role.Accounts, &gconfigv1alpha1.RoleAccount{
				ProviderId:    ra.ProviderID,
				AccountId:     ra.AccountID,
				DefaultRegion: ra.DefaultRegion,
			})
		}
		for _, rule := range r.Rules {
			role.Rules = append(role.Rules, &gconfigv1alpha1.Rule{
				Policy: rule.Policy.Policy,
				Group:  rule.Group,
				Token:  rule.RequireTicket,

				Breakglass: rule.Breakglass,
			})
		}
		out.Roles = append(out.Roles, role)
	}
	for _, t := range c.Tests {
		out.Tests = append(out.Tests, &gconfigv1alpha1.Test{
			Name: t.Name,
			Given: &gconfigv1alpha1.Given{
				User:    t.Given.User,
				Group:   t.Given.Group,
				Account: t.Given.Account,
				Role:    t.Given.Role,
			},
			Then: &gconfigv1alpha1.Then{
				Outcome: t.Then.Outcome,
			},
		})
	}

	return out
}

func FromProtobuf(c *gconfigv1alpha1.Config, providers *gconfigv1alpha1.Providers) Config {
	out := Config{
		Type:      "granted/v1alpha1",
		providers: providers,
	}

	for _, u := range c.Admins {
		out.Admins = append(out.Admins, Member{
			Email: u.Email,
		})
	}
	for _, u := range c.Users {
		out.Users = append(out.Users, Member{
			Email: u.Email,
		})
	}
	for _, g := range c.Groups {
		group := Group{
			Name: g.Name,
			ID:   g.Id,
		}
		for _, u := range g.Members {
			group.Members = append(group.Members, Member{
				Email: u.Email,
			})
		}
		out.Groups = append(out.Groups, group)
	}
	for _, r := range c.Roles {
		role := Role{
			ID:              r.Id,
			Policy:          r.Policy,
			SessionDuration: r.SessionDuration.AsDuration(),
			Group:           r.Group,
			Type:            r.Type.String(),
		}
		for _, ra := range r.Accounts {
			role.roleAccounts = append(role.roleAccounts, RoleAccount{
				ProviderID: ra.ProviderId,
				AccountID:  ra.AccountId,
			})
			role.Accounts = append(role.Accounts, Account{Account: ra.AccountId, DefaultRegion: ra.DefaultRegion})
		}
		for _, rule := range r.Rules {
			role.Rules = append(role.Rules, Rule{
				Policy:        RulePolicyField{Policy: rule.Policy},
				Group:         rule.Group,
				RequireTicket: rule.Token,

				Breakglass: rule.Breakglass,
			})
		}
		out.Roles = append(out.Roles, &role)
	}
	for _, t := range c.Tests {
		out.Tests = append(out.Tests, Test{
			Name: t.Name,
			Given: Given{
				User:    t.Given.User,
				Group:   t.Given.Group,
				Account: t.Given.Account,
				Role:    t.Given.Role,
			},
			Then: Then{
				Outcome: t.Then.Outcome,
			},
		})
	}

	return out
}
