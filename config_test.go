package gconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_LineNumberParsed(t *testing.T) {
	str := `id: dev
name: Development
provider: aws
awsAccountId: 123456789012
`

	var a Account

	err := yaml.Unmarshal([]byte(str), &a)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, a.pos.Line)
}

func Test_ChildAccountsParsed(t *testing.T) {
	str := `id: dev
name: Group
provider: aws
accounts:
  - id: one
    name: Development
    awsAccountId: 123456789012
`

	var a Account

	err := yaml.Unmarshal([]byte(str), &a)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "one", a.Children[0].ID)
}
