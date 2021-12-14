package gconfig

import (
	"testing"

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

	c, err := parseContents("config.yml", []byte(str))
	if err != nil {
		t.Fatal(err)
	}

	errs := c.Validate()
	assert.Equal(t, "config.yml:12:7: c@test.com must be defined as a user or an admin", errs.Error())
}

func TestValidAccounts(t *testing.T) {
	str := `providers:
  - id: aws
    type: awsRole
    bastionAccountId: 12345678912

accounts:
  - id: dev
    name: Development
    provider: aws
    awsAccountId: 123456789012

  - id: prod-group
    name: AWS Production
    provider: aws
    accounts:
      - id: prod-service-a
        name: Service A (production)
        awsAccountId: 123456789012
      - id: prod-service-b
        name: Service B (production)
        awsAccountId: 123456789012
`

	var c Config
	err := yaml.Unmarshal([]byte(str), &c)
	if err != nil {
		t.Fatal(err)
	}

	errs := c.Validate()
	assert.Nil(t, errs)
}

func TestDuplicateAccounts(t *testing.T) {
	str := `providers:
  - id: aws
    type: awsRole
    bastionAccountId: 12345678912

accounts:
  - id: dev
    name: Development
    provider: aws
    awsAccountId: 123456789012

  - id: dev
    name: Development
    provider: aws
    awsAccountId: 123456789012

`

	c, err := parseContents("config.yml", []byte(str))
	if err != nil {
		t.Fatal(err)
	}

	errs := c.Validate()
	assert.Equal(t, "config.yml:12:5: duplicate account ID dev", errs.Error())
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

func TestInvalidAWSAccount(t *testing.T) {
	str := `providers:
  - id: aws
    type: awsRole
    bastionAccountId: 123123
`

	c, err := parseContents("config.yml", []byte(str))
	if err != nil {
		t.Fatal(err)
	}

	errs := c.Validate()
	assert.Equal(t, "config.yml:2:5: account 123123 is not a valid AWS account: must be 12 characters long", errs.Error())
}
