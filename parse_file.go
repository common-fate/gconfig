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

	for _, r := range c.Roles {
		r.pos.Filename = filename
	}

	return &c, nil
}
