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
	for _, p := range c.Providers {
		provider := &gconfigv1alpha1.Provider{
			Id:   p.ID,
			Type: p.Type,
		}
		if p.BastionAccountID != nil {
			provider.BastionAccountId = *p.BastionAccountID
		}
		if p.InstanceARN != nil {
			provider.InstanceArn = *p.InstanceARN
		}
		if p.IdentityStoreID != nil {
			provider.IdentityStoreId = *p.IdentityStoreID
		}

		out.Providers = append(out.Providers, provider)
	}
	for _, a := range c.Accounts {
		out.Accounts = append(out.Accounts, a.SerializeProtobuf())
	}
	for _, r := range c.Roles {
		role := &gconfigv1alpha1.Role{
			Id:     r.ID,
			Policy: r.Policy,
		}
		role.Accounts = append(role.Accounts, r.Accounts...)
		for _, rule := range r.Rules {
			role.Rules = append(role.Rules, &gconfigv1alpha1.Rule{
				Policy:          rule.Policy,
				Group:           rule.Group,
				SessionDuration: durationpb.New(rule.SessionDuration),
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
				Audited: t.Then.Audited,
			},
		})
	}

	return out
}

func (a Account) SerializeProtobuf() *gconfigv1alpha1.Account {
	out := &gconfigv1alpha1.Account{
		Id:   a.ID,
		Name: a.Name,
	}
	if a.Provider != nil {
		out.Provider = *a.Provider
	}
	if a.AwsAccountID != nil {
		out.AwsAccountId = *a.AwsAccountID
	}
	for _, child := range a.Children {
		childOut := child.SerializeProtobuf()
		out.Children = append(out.Children, childOut)
	}

	return out
}

func FromProtobuf(c *gconfigv1alpha1.Config) Config {
	out := Config{
		Type: "granted/v1alpha1",
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
	for _, p := range c.Providers {
		provider := Provider{
			ID:   p.Id,
			Type: p.Type,
		}
		if p.BastionAccountId != "" {
			provider.BastionAccountID = &p.BastionAccountId
		}
		if p.InstanceArn != "" {
			provider.InstanceARN = &p.InstanceArn
		}
		if p.IdentityStoreId != "" {
			provider.IdentityStoreID = &p.IdentityStoreId
		}

		out.Providers = append(out.Providers, provider)
	}
	for _, a := range c.Accounts {
		out.Accounts = append(out.Accounts, AccountFromProtobuf(a))
	}
	for _, r := range c.Roles {
		role := Role{
			ID:     r.Id,
			Policy: r.Policy,
		}
		role.Accounts = append(role.Accounts, r.Accounts...)
		for _, rule := range r.Rules {
			role.Rules = append(role.Rules, Rule{
				Policy:          rule.Policy,
				Group:           rule.Group,
				SessionDuration: rule.SessionDuration.AsDuration(),
			})
		}
		out.Roles = append(out.Roles, role)
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
				Audited: t.Then.Audited,
			},
		})
	}

	return out
}

func AccountFromProtobuf(a *gconfigv1alpha1.Account) Account {
	out := Account{
		ID:   a.Id,
		Name: a.Name,
	}
	if a.Provider != "" {
		out.Provider = &a.Provider
	}
	if a.AwsAccountId != "" {
		out.AwsAccountID = &a.AwsAccountId
	}

	for _, child := range a.Children {
		childOut := AccountFromProtobuf(child)
		out.Children = append(out.Children, childOut)
	}

	return out
}
