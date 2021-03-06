package gconfig

import (
	"time"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"gopkg.in/yaml.v3"
)

// Config for Granted.
type Config struct {
	Type              string              `yaml:"type"`
	Admins            []Member            `yaml:"admins"`
	Users             []Member            `yaml:"users"`
	Groups            []Group             `yaml:"groups"`
	Roles             []*Role             `yaml:"roles"`
	Tests             []Test              `yaml:"tests"`
	ProviderOverrides []ProviderOverrides `yaml:"providers"`
	providers         *gconfigv1alpha1.Providers
}

// GetProviders returns the providers associated with the config.
// These providers are set when the config is parsed from YAML
func (c *Config) GetProviders() []*gconfigv1alpha1.Provider {
	if c.providers == nil {
		return nil
	}
	return c.providers.Providers
}

type ProviderOverrides struct {
	ID            string `yaml:"id"`
	DefaultRegion string `yaml:"defaultRegion"`

	// pos is used for displaying linting errors
	pos *FilePosition
}

func (p ProviderOverrides) filePosition() *FilePosition {
	return p.pos
}

func (p *ProviderOverrides) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		ID            string `yaml:"id"`
		DefaultRegion string `yaml:"defaultRegion"`
	}

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	p.ID = tmp.ID
	p.DefaultRegion = tmp.DefaultRegion

	// Save the line number
	p.pos = &FilePosition{
		Col:  value.Column,
		Line: value.Line,
	}

	return nil
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

type Account struct {
	Account       string `yaml:"acct"`
	DefaultRegion string `yaml:"defaultRegion"`

	// pos is used for displaying linting errors
	pos *FilePosition
}

func (a Account) filePosition() *FilePosition {
	return a.pos
}

func (a *Account) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		Account       string `yaml:"acct"`
		DefaultRegion string `yaml:"defaultRegion"`
	}

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	a.Account = tmp.Account
	a.DefaultRegion = tmp.DefaultRegion

	// Save the line number
	a.pos = &FilePosition{
		Col:  value.Column,
		Line: value.Line,
	}

	return nil
}

type RulePolicyField struct {
	Policy string
	// pos is used for displaying linting errors
	pos *FilePosition
}

type Rule struct {
	Policy        RulePolicyField `yaml:"policy"`
	Group         string          `yaml:"group"`
	RequireTicket bool            `yaml:"requireTicket,omitempty"`
	Breakglass    bool            `yaml:"breakglass"`
}

type Role struct {
	ID              string        `yaml:"id"`
	Accounts        []Account     `yaml:"accounts"`
	Policy          string        `yaml:"policy"`
	SessionDuration time.Duration `yaml:"sessionDuration"`
	Rules           []Rule        `yaml:"rules"`
	Group           string        `yaml:"group"`
	Type            string        `yaml:"type"`
	DefaultRegion   string        `yaml:"defaultRegion"`
	roleAccounts    []RoleAccount
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
		ID              string        `yaml:"id"`
		Accounts        []Account     `yaml:"accounts"`
		Policy          string        `yaml:"policy"`
		Rules           []Rule        `yaml:"rules"`
		SessionDuration time.Duration `yaml:"sessionDuration"`
		Group           string        `yaml:"group"`
		Type            string        `yaml:"type"`
		DefaultRegion   string        `yaml:"defaultRegion"`
	}

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	r.ID = tmp.ID
	r.Accounts = tmp.Accounts
	r.Policy = tmp.Policy
	r.Rules = tmp.Rules
	r.SessionDuration = tmp.SessionDuration
	r.Type = tmp.Type
	r.Group = tmp.Group
	r.DefaultRegion = tmp.DefaultRegion

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

// RoleAccount is a binding of a role to an Account
// in a particular provider
type RoleAccount struct {
	AccountID     string
	ProviderID    string
	DefaultRegion string
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
}

type FilePosition struct {
	Filename string
	Col      int
	Line     int
}
