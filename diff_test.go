package gconfig

import (
	"testing"

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

func TestEmpty(t *testing.T) {
	c := Changes{}
	assert.Equal(t, true, c.Empty())
}
