package gconfig

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// ParseFile parses a Granted YAML config file
func ParseFile(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContents(filename, b)
}

// parseContents parses a Granted YAML config file
func parseContents(filename string, in []byte) (*Config, error) {
	var c Config

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

	for _, p := range c.Providers {
		p.pos.Filename = filename
	}

	for _, a := range c.Accounts {
		setAccountFilename(a, filename)
	}

	return &c, nil
}

func setAccountFilename(a Account, filename string) {
	a.pos.Filename = filename
	for _, c := range a.Children {
		setAccountFilename(c, filename)
	}
}
