package gcoktav1alpha1

import (
	pbgcoktav1alpha1 "github.com/common-fate/gconfig/gen/gconfig/okta/v1alpha1"
)

func (c *Config) SerializeProtobuf() *pbgcoktav1alpha1.Config {
	out := &pbgcoktav1alpha1.Config{}

	for _, u := range c.Admins {
		out.Admins = append(out.Admins, &pbgcoktav1alpha1.Member{
			Email: u.Email,
		})
	}

	for _, r := range c.Roles {
		role := &pbgcoktav1alpha1.Role{
			Id: r.ID,
		}

		for _, rule := range r.Rules {
			role.Rules = append(role.Rules, &pbgcoktav1alpha1.Rule{
				Policy: rule.Policy.Policy,
				Group:  rule.Group,

				Breakglass: rule.Breakglass,
			})
		}
		out.Roles = append(out.Roles, role)
	}

	return out
}

func FromProtobuf(c *pbgcoktav1alpha1.Config) Config {
	out := Config{
		Type: "granted/v1alpha1",
	}

	for _, u := range c.Admins {
		out.Admins = append(out.Admins, Member{
			Email: u.Email,
		})
	}

	for _, r := range c.Roles {
		role := Role{
			ID: r.Id,
		}

		for _, rule := range r.Rules {
			role.Rules = append(role.Rules, Rule{
				Policy: RulePolicyField{Policy: rule.Policy},
				Group:  rule.Group,

				Breakglass: rule.Breakglass,
			})
		}
		out.Roles = append(out.Roles, &role)
	}

	return out
}
