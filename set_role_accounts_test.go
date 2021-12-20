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
