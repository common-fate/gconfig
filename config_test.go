package gconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_LineNumberParsed(t *testing.T) {
	str := `admins:
  - user@test.com
`

	var c Config

	err := yaml.Unmarshal([]byte(str), &c)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, c.Admins)
	assert.Equal(t, 2, c.Admins[0].pos.Line)
}
