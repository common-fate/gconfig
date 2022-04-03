package gconfig

import (
	"testing"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestSetRoleAccounts(t *testing.T) {
	str := `roles:
  - id: test
    type: aws
    accounts: 
      - acct:  "123456789012"
    policy: TEST_POLICY
    sessionDuration: 8h
  `

	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:   "123456789012",
					},
				}}},
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

func TestSetRoleRuleAccounts(t *testing.T) {
	str := `roles:
  - id: test
    type: aws
    accounts: 
      - acct:  "123456789012"
    policy: TEST_POLICY
    sessionDuration: 8h
    rules:
      - policy: 
          allow: true
        group: developers
        
  `

	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:   "123456789012",
					},
				}}},
			},
		},
	}

	c, err := parseContents("config.yml", []byte(str), providers)
	if err != nil {
		t.Fatal(err)
	}

	expected := Rule{
		Group:  "developers",
		Policy: map[string]interface{}{"allow": true},
		pos:    c.Roles[0].Rules[0].pos,
	}

	actual := c.Roles[0].Rules[0]

	assert.Equal(t, expected, actual)
}

func TestSetRoleAccounts_Invalid(t *testing.T) {
	str := `roles:
  - id: test
    type: aws
    accounts: 
      - acct:  "123456789012"
    policy: TEST_POLICY
    sessionDuration: 8h
  `

	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:   "222333444555",
					},
				}}},
			},
		},
	}

	_, err := parseContents("config.yml", []byte(str), providers)
	assert.Equal(t, "config.yml:2:5: role test references an account that doesn't exist: 123456789012", err.Error())
}

func TestSetRoleAccounts_Alias(t *testing.T) {
	str := `roles:
  - id: test
    type: aws
    accounts: 
      - acct:  "dev"
    policy: TEST_POLICY
    sessionDuration: 8h
  `

	// an alias "dev" is provided for account ID 123456789012
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:      "123456789012",
						Aliases: []string{"dev"},
					},
				}}},
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
    type: aws
    accounts: 
      - acct:  "Develop"
    policy: TEST_POLICY
    sessionDuration: 8h
  `

	// a name "Develop" is provided for account ID 123456789012
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:   "123456789012",
						Name: "Develop",
					},
				}}},
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
    type: aws
    accounts: 
      - acct:  "dev"
    policy: TEST_POLICY
    sessionDuration: 8h
  `

	// multiple accounts have the same "dev" alias. In this case the admin is not allowed to
	// use the "dev" alias and must provide the fully-qualified account ID (e.g. aws:123456789012)
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
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
				}}},
			},
		},
	}

	_, err := parseContents("config.yml", []byte(str), providers)
	// TODO: line and character number will need to be updated once feature is implemented
	assert.Equal(t, "config.yml:2:5: role test: account 'dev' is ambiguous and could refer to one of these account names:\n\n    - aws:dev:123456789012 (TYPE_AWS_ACCOUNT 123456789012 in provider aws)\n    - aws:dev:222333444555 (TYPE_AWS_ACCOUNT 222333444555 in provider aws)\n\nPlease replace 'dev' with the account name above that you meant (e.g. aws:dev:123456789012).", err.Error())
}

func TestSetRoleAccounts_PartialAlias(t *testing.T) {
	str := `roles:
  - id: test
    type: aws
    accounts: 
      - acct:  "aws:dev"
    policy: TEST_POLICY
    sessionDuration: 8h	
  `

	// an alias "dev" is provided for account ID 123456789012
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:      "123456789012",
						Aliases: []string{"dev"},
					},
				}}},
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

func TestSetRoleAccounts_FullWithAlias(t *testing.T) {
	str := `roles:
  - id: test
    type: aws
    accounts: 
      - acct:  "aws:dev:123456789012"
    policy: TEST_POLICY
    sessionDuration: 8h	
  `

	// an alias "dev" is provided for account ID 123456789012
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
						Id:      "123456789012",
						Aliases: []string{"dev"},
					},
				}}},
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

func TestSetRoleAccounts_OU(t *testing.T) {
	str := `roles:
  - id: test
    type: aws
    accounts: 
      - acct:  "ou-4w0n-bads234"
    policy: TEST_POLICY
    sessionDuration: 8h	
  `

	// a name "dev" is provided for account ID 123456789012 and the ou "ou-4w0n-bads234"
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type:     gconfigv1alpha1.Account_TYPE_UNSPECIFIED,
						Id:       "ou-4w0n-bads234",
						Name:     "dev",
						Children: []*gconfigv1alpha1.Account{{Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT, Id: "12345678912", Name: "dev"}, {Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT, Id: "02345678912", Name: "admin"}},
					},
				}}},
			},
		},
	}

	c, err := parseContents("config.yml", []byte(str), providers)
	if err != nil {
		t.Fatal(err)
	}

	// the RoleAccount should still refer to the account ID rather than the alias, there should be a single role account for the ou
	expected1 := RoleAccount{
		AccountID:  "12345678912",
		ProviderID: "aws",
	}
	expected2 := RoleAccount{
		AccountID:  "02345678912",
		ProviderID: "aws",
	}
	if len(c.Roles) != 1 {
		t.Fatal("expected 1 role to be parsed")
	}
	if len(c.Roles[0].roleAccounts) != 2 {
		t.Fatal("expected 2 roleaccounts to be generated")
	}

	actual := c.Roles[0].roleAccounts[0]

	assert.Equal(t, expected1, actual)

	actual = c.Roles[0].roleAccounts[1]

	assert.Equal(t, expected2, actual)
}

func TestSetRoleAccounts_FullWithAliasOU(t *testing.T) {
	str := `roles:
  - id: test
    type: aws
    accounts: 
      - acct:  "aws:dev:ou-4w0n-bads234"
    policy: TEST_POLICY
    sessionDuration: 8h	
  `

	// a name "dev" is provided for account ID 123456789012 and the ou "ou-4w0n-bads234"
	providers := &gconfigv1alpha1.Providers{
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id: "aws",
				Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
					{
						Type:     gconfigv1alpha1.Account_TYPE_UNSPECIFIED,
						Id:       "ou-4w0n-bads234",
						Name:     "dev",
						Children: []*gconfigv1alpha1.Account{{Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT, Id: "12345678912", Name: "dev"}, {Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT, Id: "02345678912", Name: "admin"}},
					},
				}}},
			},
		},
	}

	c, err := parseContents("config.yml", []byte(str), providers)
	if err != nil {
		t.Fatal(err)
	}

	// the RoleAccount should still refer to the account ID rather than the alias, there should be a single role account for the ou
	expected1 := RoleAccount{
		AccountID:  "12345678912",
		ProviderID: "aws",
	}
	expected2 := RoleAccount{
		AccountID:  "02345678912",
		ProviderID: "aws",
	}
	if len(c.Roles) != 1 {
		t.Fatal("expected 1 role to be parsed")
	}
	if len(c.Roles[0].roleAccounts) != 2 {
		t.Fatal("expected 2 roleaccounts to be generated")
	}

	actual := c.Roles[0].roleAccounts[0]

	assert.Equal(t, expected1, actual)

	actual = c.Roles[0].roleAccounts[1]

	assert.Equal(t, expected2, actual)
}
