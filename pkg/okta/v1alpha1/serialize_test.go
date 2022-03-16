package gcoktav1alpha1

import (
	"testing"

	pbgcoktav1alpha1 "github.com/common-fate/gconfig/gen/gconfig/okta/v1alpha1"
	"github.com/stretchr/testify/assert"
)

var (
	cfg = Config{

		Type: "granted/v1alpha1",
		Admins: []Member{
			{
				Email: "admin@example.com",
			},
		},

		Roles: []*Role{
			{
				ID: "role",

				Rules: []Rule{
					{
						Policy: RulePolicyField{Policy: RulePolicyAllow.String()},
						Group:  "test",
					},
				},
			},
		},
	}

	expected = &pbgcoktav1alpha1.Config{
		Admins: []*pbgcoktav1alpha1.Member{
			{
				Email: "admin@example.com",
			},
		},

		Roles: []*pbgcoktav1alpha1.Role{
			{
				Id: "role",

				Rules: []*pbgcoktav1alpha1.Rule{
					{
						Policy: "allow",
						Group:  "test",
					},
				},
			},
		},
	}
)

func TestSerialize(t *testing.T) {
	out := cfg.SerializeProtobuf()
	assert.Equal(t, expected, out)
}

func TestDeserialize(t *testing.T) {
	reversed := FromProtobuf(expected)
	assert.Equal(t, cfg, reversed)
}
