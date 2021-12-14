package gconfig

import (
	"time"

	"gopkg.in/yaml.v3"
)

// Config for Granted.
type Config struct {
	Type      string     `yaml:"type"`
	Admins    []Member   `yaml:"admins"`
	Users     []Member   `yaml:"users"`
	Groups    []Group    `yaml:"groups"`
	Providers []Provider `yaml:"providers"`
	Accounts  []Account  `yaml:"accounts"`
	Roles     []Roles    `yaml:"roles"`
	Tests     []Tests    `yaml:"tests"`
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

type Provider struct {
	ID               string  `yaml:"id"`
	Type             string  `yaml:"type"`
	BastionAccountID *string `yaml:"bastionAccountId,omitempty"`
	InstanceARN      *string `yaml:"instanceARN,omitempty"`
	IdentityStoreID  *string `yaml:"identityStoreId,omitempty"`

	// pos is used for displaying linting errors
	pos *FilePosition
}

func (p *Provider) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		ID               string  `yaml:"id"`
		Type             string  `yaml:"type"`
		BastionAccountID *string `yaml:"bastionAccountId,omitempty"`
		InstanceARN      *string `yaml:"instanceARN,omitempty"`
		IdentityStoreID  *string `yaml:"identityStoreId,omitempty"`
	}

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	p.ID = tmp.ID
	p.Type = tmp.Type
	p.BastionAccountID = tmp.BastionAccountID
	p.InstanceARN = tmp.InstanceARN
	p.IdentityStoreID = tmp.IdentityStoreID

	// Save the line number
	p.pos = &FilePosition{
		Col:  value.Column,
		Line: value.Line,
	}

	return nil
}

func (p Provider) filePosition() *FilePosition {
	return p.pos
}

type Account struct {
	ID           string    `yaml:"id"`
	Name         string    `yaml:"name"`
	Provider     *string   `yaml:"provider,omitempty"`
	AwsAccountID *string   `yaml:"awsAccountId,omitempty"`
	Children     []Account `yaml:"accounts,omitempty"`

	// the ID of the parent account if it has one
	parentId *string

	// pos is used for displaying linting errors
	pos *FilePosition
}

func (a *Account) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		ID           string      `yaml:"id"`
		Name         string      `yaml:"name"`
		Provider     *string     `yaml:"provider,omitempty"`
		AwsAccountID *string     `yaml:"awsAccountId,omitempty"`
		Children     []yaml.Node `yaml:"accounts,omitempty"`
	}

	err := value.Decode(&tmp)
	if err != nil {
		return err
	}

	a.ID = tmp.ID
	a.Name = tmp.Name
	a.Provider = tmp.Provider
	a.AwsAccountID = tmp.AwsAccountID

	for _, c := range tmp.Children {
		var childAcc Account
		err := c.Decode(&childAcc)
		if err != nil {
			return err
		}
		childAcc.parentId = &a.ID
		a.Children = append(a.Children, childAcc)
	}

	// Save the line number
	a.pos = &FilePosition{
		Col:  value.Column,
		Line: value.Line,
	}

	return nil
}

func (a *Account) filePosition() *FilePosition {
	return a.pos
}

type Rules struct {
	Policy          string        `yaml:"policy"`
	Group           string        `yaml:"group"`
	SessionDuration time.Duration `yaml:"sessionDuration"`
}

type Roles struct {
	ID       string   `yaml:"id"`
	Accounts []string `yaml:"accounts"`
	Policy   string   `yaml:"policy"`
	Rules    []Rules  `yaml:"rules"`
}

// Tests is the container for all Granted configuration tests
// expressed as part of a Granted config.
type Tests struct {
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
