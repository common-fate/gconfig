package gcoktav1alpha1

import (
	"gopkg.in/yaml.v3"
)

// Config for Granted.
type Config struct {
	Type   string   `yaml:"type"`
	Admins []Member `yaml:"admins"`
	Roles  []*Role  `yaml:"roles"`
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

type RulePolicyField struct {
	Policy string
	// pos is used for displaying linting errors
	pos *FilePosition
}

type Rule struct {
	Policy     RulePolicyField `yaml:"policy"`
	Group      string          `yaml:"group"`
	Breakglass bool            `yaml:"breakglass"`
}

type Role struct {
	ID    string `yaml:"id"`
	Group string `yaml:"group"`
	Rules []Rule `yaml:"rules"`
	// pos is used for displaying linting errors
	pos *FilePosition
}

// validates that the policy is valid at parsing time
// To add more policy types, add to the policy.go enum
func (p *RulePolicyField) UnmarshalYAML(value *yaml.Node) error {
	var tmp string
	err := value.Decode(&tmp)
	if err != nil {
		return err
	}
	p.pos = &FilePosition{
		Col:  value.Column,
		Line: value.Line,
	}
	p.Policy = tmp
	return nil

}
func (p *RulePolicyField) filePosition() *FilePosition {
	return p.pos
}
func (r *Role) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		ID    string `yaml:"id"`
		Rules []Rule `yaml:"rules"`
		Group string `yaml:"group"`
	}

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	r.ID = tmp.ID
	r.Rules = tmp.Rules
	r.Group = tmp.Group

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

type FilePosition struct {
	Filename string
	Col      int
	Line     int
}
