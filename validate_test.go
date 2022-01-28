package gconfig

import (
	"testing"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestValidateGroupMembers(t *testing.T) {
	str := `admins:
- a@test.com

users:
- b@test.com

groups:
- name: AWS Developers
  id: developers
  members:
    - a@test.com
    - b@test.com
`

	var c Config
	err := yaml.Unmarshal([]byte(str), &c)
	if err != nil {
		t.Fatal(err)
	}

	errs := c.Validate()
	assert.Nil(t, errs)
}

// c@test.com hasn't been defined as a user or admin.
func TestValidateGroupMembersInvalid(t *testing.T) {
	str := `admins:
- a@test.com

users:
- b@test.com

groups:
- name: AWS Developers
  id: developers
  members:
    - a@test.com
    - c@test.com
`

	c, err := parseContents("config.yml", []byte(str), &gconfigv1alpha1.Providers{})
	if err != nil {
		t.Fatal(err)
	}

	errs := c.Validate()
	assert.Equal(t, "config.yml:12:7: c@test.com must be defined as a user or an admin", errs.Error())
}

// If we construct a Config using Go structs,
// we should still be able to validate it and the resulting
// validation errors shouldn't contain filenames or line numbers.
func TestErrorPrintingNoFilename(t *testing.T) {
	c := Config{
		Groups: []Group{
			{
				Name: "test",
				ID:   "test",
			},
			{
				Name: "test",
				ID:   "test",
			},
		},
	}

	errs := c.Validate()
	assert.Equal(t, "duplicate group ID test", errs.Error())
}

func TestValidAccounts(t *testing.T) {
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

	err = c.Validate()
	assert.NoError(t, err)
}

func TestGrantedAdministratorsGroupCannotBeCreated(t *testing.T) {
	str := `admins:
- a@test.com

groups:
- name: granted:administrators
  id: gadmins
  members:
    - a@test.com
`

	c, err := parseContents("config.yml", []byte(str), &gconfigv1alpha1.Providers{})
	if err != nil {
		t.Fatal(err)
	}

	errs := c.Validate()
	assert.Equal(t, "a group called `granted:administrators` cannot be created. Please choose a different name", errs.Error())
}

