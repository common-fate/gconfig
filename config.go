package gconfig

import (
	"time"

	"gopkg.in/yaml.v3"
)

// Config for Granted.
type Config struct {
	Type   string   `yaml:"type"`
	Admins []Member `yaml:"admins"`
	Users  []Member `yaml:"users"`
	Groups []Group  `yaml:"groups"`
	Roles  []Role   `yaml:"roles"`
	Tests  []Test   `yaml:"tests"`
}

type Group struct {
	Name    string   `yaml:"name"`
	ID      string   `yaml:"id"`
	Members []Member `yaml:"members"`

	// pos is used for displaying linting errors
	pos *FilePosition
}

func (g Group) filePosition() *FilePosition {
	return g.pos
}

func (g *Group) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		Name    string   `yaml:"name"`
		ID      string   `yaml:"id"`
		Members []Member `yaml:"members"`
	}

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	g.ID = tmp.ID
	g.Name = tmp.Name
	g.Members = tmp.Members

	// Save the line number
	g.pos = &FilePosition{
		Col:  value.Column,
		Line: value.Line,
	}

	return nil
}

type Member struct {
	Email string

	// pos is used for displaying linting errors
	pos *FilePosition
}

func (m *Member) UnmarshalYAML(value *yaml.Node) error {
	var tmp string

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	m.Email = tmp

	// Save the line number
	m.pos = &FilePosition{
		Col:  value.Column,
		Line: value.Line,
	}

	return nil
}

func (m Member) filePosition() *FilePosition {
	return m.pos
}

type Rule struct {
	Policy          string        `yaml:"policy"`
	Group           string        `yaml:"group"`
	SessionDuration time.Duration `yaml:"sessionDuration"`
}

type Role struct {
	ID       string   `yaml:"id"`
	Accounts []string `yaml:"accounts"`
	Policy   string   `yaml:"policy"`
	Rules    []Rule   `yaml:"rules"`

	// pos is used for displaying linting errors
	pos *FilePosition
}

func (r *Role) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		ID       string   `yaml:"id"`
		Accounts []string `yaml:"accounts"`
		Policy   string   `yaml:"policy"`
		Rules    []Rule   `yaml:"rules"`
	}

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	r.ID = tmp.ID
	r.Accounts = tmp.Accounts
	r.Policy = tmp.Policy

	// Save the line number
	r.pos = &FilePosition{
		Col:  value.Column,
		Line: value.Line,
	}

	return nil
}

func (r Role) filePosition() *FilePosition {
	return r.pos
}

// Test is the container for all Granted configuration tests
// expressed as part of a Granted config.
type Test struct {
	Name  string `yaml:"name"`
	Given Given  `yaml:"given,omitempty"`
	Then  Then   `yaml:"then,omitempty"`
}

type Given struct {
	User    string `yaml:"user"`
	Group   string `yaml:"group"`
	Account string `yaml:"account"`
	Role    string `yaml:"role"`
}

type Then struct {
	Outcome string `yaml:"outcome"`
	Audited *bool  `yaml:"audited,omitempty"`
}

type FilePosition struct {
	Filename string
	Col      int
	Line     int
}
