package gconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	old := Config{
		Providers: []Provider{
			{
				ID:   "first",
				Type: "first",
			},
			{
				ID:   "common",
				Type: "common",
			},
		},
	}

	new := Config{
		Providers: []Provider{
			{
				ID:   "second",
				Type: "second",
			},
			{
				ID:   "common",
				Type: "common",
			},
		},
	}

	res, err := new.ChangesFrom(old)
	if err != nil {
		t.Fatal(err)
	}

	expected := Changes{
		DeleteProviders: []string{"first"},
		AddProviders: []Provider{
			{
				ID:   "second",
				Type: "second",
			},
		},
	}

	assert.Equal(t, expected, res)
}

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

// Currently we don't support in-place updates for resoureces like
// providers and accounts - instead we return an error so the user
// knows they need to create a new provider with a different ID.
func TestUpdateProviderDiff(t *testing.T) {
	old := Config{
		Providers: []Provider{
			{
				ID:   "first",
				Type: "first",
			},
		},
	}

	new := Config{
		Providers: []Provider{
			{
				ID:   "first",
				Type: "second",
			},
		},
	}

	_, err := new.ChangesFrom(old)
	assert.Equal(t, &ErrNoInPlaceUpdates{Type: "provider", ID: "first", Field: "type", Old: "first", New: "second"}, err)
}

func TestUpdateProviderDiffNoError(t *testing.T) {
	accID := "123456"
	old := Config{
		Providers: []Provider{
			{
				ID:               "first",
				Type:             "first",
				BastionAccountID: &accID,
			},
		},
	}

	secondAccID := "123456"

	new := Config{
		Providers: []Provider{
			{
				ID:               "first",
				Type:             "first",
				BastionAccountID: &secondAccID,
			},
		},
	}

	res, err := new.ChangesFrom(old)
	if err != nil {
		t.Fatal(err)
	}

	expected := Changes{}

	assert.Equal(t, expected, res)
}

func TestEmpty(t *testing.T) {
	c := Changes{}
	assert.Equal(t, true, c.Empty())
}
