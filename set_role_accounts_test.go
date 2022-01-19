package gconfig

import (
	"testing"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestSetRoleAccounts(t *testing.T) {
	str := `roles:
  - id: test
    accounts: 
      - "123456789012"
    policy: TEST_POLICY
  `

	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Accounts: []*gconfigv1alpha1.Account{
					{
						Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:   "123456789012",
					},
				},
			},
		},
	}

	c, err := parseContents("config.yml", []byte(str), providers)
	if err != nil {
		t.Fatal(err)
	}

	expected := RoleAccount{
		AccountID:  "123456789012",
		ProviderID: "aws",
	}
	if len(c.Roles) != 1 {
		t.Fatal("expected 1 role to be parsed")
	}

	actual := c.Roles[0].roleAccounts[0]

	assert.Equal(t, expected, actual)
}

func TestSetRoleAccounts_Invalid(t *testing.T) {
	str := `roles:
  - id: test
    accounts: 
      - "123456789012"
    policy: TEST_POLICY
  `

	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Accounts: []*gconfigv1alpha1.Account{
					{
						Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:   "222333444555",
					},
				},
			},
		},
	}

	_, err := parseContents("config.yml", []byte(str), providers)
	assert.Equal(t, "config.yml:2:5: role test references an account that doesn't exist: 123456789012", err.Error())
}

func TestSetRoleAccounts_Alias(t *testing.T) {
	str := `roles:
  - id: test
    accounts: 
      - "dev"
    policy: TEST_POLICY
  `

	// an alias "dev" is provided for account ID 123456789012
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Accounts: []*gconfigv1alpha1.Account{
					{
						Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:      "123456789012",
						Aliases: []string{"dev"},
					},
				},
			},
		},
	}

	c, err := parseContents("config.yml", []byte(str), providers)
	if err != nil {
		t.Fatal(err)
	}

	// the RoleAccount should still refer to the account ID rather than the alias
	expected := RoleAccount{
		AccountID:  "123456789012",
		ProviderID: "aws",
	}
	if len(c.Roles) != 1 {
		t.Fatal("expected 1 role to be parsed")
	}

	actual := c.Roles[0].roleAccounts[0]

	assert.Equal(t, expected, actual)
}

func TestSetRoleAccounts_Name(t *testing.T) {
	str := `roles:
  - id: test
    accounts: 
      - "Develop"
    policy: TEST_POLICY
  `

	// a name "Develop" is provided for account ID 123456789012
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Accounts: []*gconfigv1alpha1.Account{
					{
						Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:   "123456789012",
						Name: "Develop",
					},
				},
			},
		},
	}

	c, err := parseContents("config.yml", []byte(str), providers)
	if err != nil {
		t.Fatal(err)
	}

	// the RoleAccount should still refer to the account ID rather than the alias
	expected := RoleAccount{
		AccountID:  "123456789012",
		ProviderID: "aws",
	}
	if len(c.Roles) != 1 {
		t.Fatal("expected 1 role to be parsed")
	}

	actual := c.Roles[0].roleAccounts[0]

	assert.Equal(t, expected, actual)
}

func TestSetRoleAccounts_ConflictingAliases(t *testing.T) {
	str := `roles:
  - id: test
    accounts: 
      - "dev"
    policy: TEST_POLICY
  `

	// multiple accounts have the same "dev" alias. In this case the admin is not allowed to
	// use the "dev" alias and must provide the fully-qualified account ID (e.g. aws:123456789012)
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Accounts: []*gconfigv1alpha1.Account{
					{
						Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:      "123456789012",
						Aliases: []string{"dev"},
					},
					{
						Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:      "222333444555",
						Aliases: []string{"dev"},
					},
				},
			},
		},
	}

	_, err := parseContents("config.yml", []byte(str), providers)
	// TODO: line and character number will need to be updated once feature is implemented
	assert.Equal(t, "config.yml:2:5: role test: account 'dev' is ambiguous and could refer to one of these account names:\n\n    - aws:dev:123456789012 (TYPE_AWS_ACCOUNT 123456789012 in provider aws)\n    - aws:dev:222333444555 (TYPE_AWS_ACCOUNT 222333444555 in provider aws)\n\nPlease replace 'dev' with the account name above that you meant (e.g. aws:dev:123456789012).", err.Error())
}
