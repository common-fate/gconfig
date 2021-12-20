package gconfig

import (
	"io/ioutil"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"gopkg.in/yaml.v3"
)

// ParseFile parses a Granted YAML config file
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
	}

	err = c.setRoleAccounts()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
