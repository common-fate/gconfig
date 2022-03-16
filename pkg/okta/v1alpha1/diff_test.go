package gcoktav1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that checks functionality for add/delete
func TestRoleDiff(t *testing.T) {
	old := Config{
		Roles: []*Role{
			{
				ID: "dev",
			},
		},
	}

	new := Config{
		Roles: []*Role{
			{
				ID: "admin",
			},
		},
	}

	res, err := new.ChangesFrom(old)
	if err != nil {
		t.Fatal(err)
	}

	expected := Changes{
		DeleteRoles: []string{"dev"},
		AddRoles:    []string{"admin"},
	}

	assert.Equal(t, expected, res)
}

// Test that checks, create, update, and delete functionality for Roles
func TestUpdateRoleDiff(t *testing.T) {
	old := Config{
		Roles: []*Role{
			{
				ID: "admin",

				Rules: []Rule{
					{
						Policy: RulePolicyField{Policy: RulePolicyAllow.String()},
					},
				},
			},
			{
				ID: "dev"},

			{
				ID: "dev2",
			},
		},
	}

	// update admin role: remove rule, and account
	// delete dev2
	// create new role for dev3
	new := Config{
		Roles: []*Role{
			{
				ID: "admin",
			},
			{
				ID: "dev",
			},
			{
				ID: "dev3",
			},
		},
	}

	res, err := new.ChangesFrom(old)
	if err != nil {
		t.Fatal(err)
	}

	expected := Changes{
		UpdateRoles: []UpdateRole{
			{
				ID:           "admin",
				AlteredField: []string{"Rules", "Rules"},

				AddRules: nil,

				DeleteRules: []DeleteRule{{Group: "", Policy: "allow", Breakglass: false}},
			},
		},
		DeleteRoles: []string{"dev2"},
		AddRoles:    []string{"dev3"},
	}

	assert.Equal(t, expected, res)
}

func TestEmpty(t *testing.T) {
	c := Changes{}
	assert.Equal(t, true, c.Empty())
}
