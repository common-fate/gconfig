package gconfig

import (
	"fmt"
	"io/ioutil"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"gopkg.in/yaml.v3"
)

// ParseFile parses a Granted YAML config file and implements read time checks
// See validate.go for any checks on the config post reading the file
func ParseFile(filename string, providers *gconfigv1alpha1.Providers) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContents(filename, b, providers)
}

// parseContents parses a Granted YAML config file
func parseContents(filename string, in []byte, providers *gconfigv1alpha1.Providers) (*Config, error) {
	var c Config

	c.providers = providers

	err := yaml.Unmarshal(in, &c)
	if err != nil {
		return nil, err
	}

	// add the filename to entities in the config
	for _, a := range c.Admins {
		a.pos.Filename = filename
	}

	for _, u := range c.Users {
		u.pos.Filename = filename
	}

	for _, g := range c.Groups {
		g.pos.Filename = filename

		for _, m := range g.Members {
			m.pos.Filename = filename
		}
	}

	for _, r := range c.Roles {
		r.pos.Filename = filename
		for _, rule := range r.Rules {
			rule.Policy.pos.Filename = filename
			// Validates that the rule policy matches a supported policy type
			if _, err := RulePolicyString(rule.Policy.Policy); err != nil {
				policyValues := []string{}
				for _, pol := range RulePolicyValues() {
					policyValues = append(policyValues, pol.String())
				}
				err = fmt.Errorf("policy: %s must be one of %v", rule.Policy.Policy, policyValues)
				err = printLintError(&rule.Policy, err)
				return nil, err
			}

		}
	}

	err = c.setRoleAccounts()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
