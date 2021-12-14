package gconfig

import (
	configv1alpha1 "github.com/common-fate/gconfig/gen/proto/go/config/v1alpha1"
	"google.golang.org/protobuf/types/known/durationpb"
)

func (c *Config) SerializeProtobuf() *configv1alpha1.Config {
	out := &configv1alpha1.Config{}

	for _, u := range c.Admins {
		out.Admins = append(out.Admins, &configv1alpha1.Member{
			Email: u.Email,
		})
	}
	for _, u := range c.Users {
		out.Users = append(out.Users, &configv1alpha1.Member{
			Email: u.Email,
		})
	}
	for _, g := range c.Groups {
		group := &configv1alpha1.Group{
			Name: g.Name,
			Id:   g.ID,
		}
		for _, u := range g.Members {
			group.Members = append(group.Members, &configv1alpha1.Member{
				Email: u.Email,
			})
		}
		out.Groups = append(out.Groups, group)
	}
	for _, p := range c.Providers {
		provider := &configv1alpha1.Provider{
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
		role := &configv1alpha1.Role{
			Id:     r.ID,
			Policy: r.Policy,
		}
		role.Accounts = append(role.Accounts, r.Accounts...)
		for _, rule := range r.Rules {
			role.Rules = append(role.Rules, &configv1alpha1.Rule{
				Policy:          rule.Policy,
				Group:           rule.Group,
				SessionDuration: durationpb.New(rule.SessionDuration),
			})
		}
		out.Roles = append(out.Roles, role)
	}
	for _, t := range c.Tests {
		out.Tests = append(out.Tests, &configv1alpha1.Test{
			Name: t.Name,
			Given: &configv1alpha1.Given{
				User:    t.Given.User,
				Group:   t.Given.Group,
				Account: t.Given.Account,
				Role:    t.Given.Role,
			},
			Then: &configv1alpha1.Then{
				Outcome: t.Then.Outcome,
				Audited: t.Then.Audited,
			},
		})
	}

	return out
}

func (a Account) SerializeProtobuf() *configv1alpha1.Account {
	out := &configv1alpha1.Account{
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
