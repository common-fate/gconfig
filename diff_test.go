package gconfig

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserDiff(t *testing.T) {
	old := Config{
		Users: []Member{
			{
				Email: "common@test.com",
			},
			{
				Email: "old@test.com",
			},
		},
	}

	new := Config{
		Users: []Member{
			{
				Email: "common@test.com",
			},
			{
				Email: "new@test.com",
			},
		},
	}

	res, err := new.ChangesFrom(old)
	if err != nil {
		t.Fatal(err)
	}

	expected := Changes{
		DeleteUsers: []string{"old@test.com"},
		AddUsers:    []string{"new@test.com"},
	}

	assert.Equal(t, expected, res)
}

func TestUpdateUserDiff(t *testing.T) {
	old := Config{
		Users: []Member{
			{
				Email: "common@test.com",
			},
			{
				Email: "old@test.com",
			},
		},
	}

	new := Config{
		Admins: []Member{
			{
				Email: "old@test.com",
			},
		},
		Users: []Member{
			{
				Email: "common@test.com",
			},
		},
	}

	res, err := new.ChangesFrom(old)
	if err != nil {
		t.Fatal(err)
	}

	expected := Changes{
		UpdateUsers: []UpdateUser{
			{
				Email: "old@test.com",
				Admin: true,
			},
		},
	}

	assert.Equal(t, expected, res)
}

// Test that checks functionality for add/delete
func TestRoleDiff(t *testing.T) {
	old := Config{
		Roles: []*Role{
			{
				ID:       "dev",
				Accounts: []Account{{Account: "123456789012"}},
			},
		},
	}

	new := Config{
		Roles: []*Role{
			{
				ID:       "admin",
				Accounts: []Account{{Account: "123456789012"}},
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
				ID:       "admin",
				Accounts: []Account{{Account: "123456789012"}, {Account: "123456789013"}},
				Rules: []Rule{
					{
						Policy: RulePolicyField{Policy: RulePolicyAllow.String()},
					},
				},
				SessionDuration: time.Hour,
			},
			{
				ID:              "dev",
				Accounts:        []Account{{Account: "123456789012"}, {Account: "123456789013"}},
				SessionDuration: time.Hour,
			},
			{
				ID:              "dev2",
				Accounts:        []Account{{Account: "123456789012"}, {Account: "123456789013"}},
				SessionDuration: time.Hour,
			},
		},
	}

	// update admin role: remove rule, and account
	// delete dev2
	// create new role for dev3
	new := Config{
		Roles: []*Role{
			{
				ID:              "admin",
				Accounts:        []Account{{Account: "123456789013"}},
				SessionDuration: time.Hour,
			},
			{
				ID:              "dev",
				Accounts:        []Account{{Account: "123456789012"}, {Account: "123456789013"}},
				SessionDuration: time.Hour,
			},
			{
				ID:              "dev3",
				Accounts:        []Account{{Account: "123456789012"}, {Account: "123456789013"}},
				SessionDuration: time.Hour,
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
				AlteredField: []string{"Rules", "Rules", "Accounts"},

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
