package gconfig

import (
	"testing"
	"time"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/proto/go/gconfig/v1alpha1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestSerialize(t *testing.T) {
	provider := "provider"
	accID := "123456789012"
	cfg := Config{
		Type: "v1alpha1",
		Admins: []Member{
			{
				Email: "admin@example.com",
			},
		},
		Users: []Member{
			{
				Email: "user@example.com",
			},
		},
		Groups: []Group{
			{
				Name: "group",
				ID:   "group",
				Members: []Member{
					{
						Email: "user@example.com",
					},
				},
			},
		},
		Providers: []Provider{
			{
				ID:               provider,
				Type:             "awsRole",
				BastionAccountID: &accID,
			},
		},
		Accounts: []Account{
			{
				ID:       "accgroup",
				Name:     "account group",
				Provider: &provider,
				Children: []Account{
					{
						ID:           "acc",
						Name:         "account",
						AwsAccountID: &accID,
					},
				},
			},
		},
		Roles: []Roles{
			{
				ID:       "role",
				Accounts: []string{"acc"},
				Policy:   "policy",
				Rules: []Rules{
					{
						Policy:          "allow",
						Group:           "test",
						SessionDuration: time.Hour,
					},
				},
			},
		},
		Tests: []Tests{
			{
				Name: "test",
				Given: Given{
					User:    "test@example.com",
					Account: "acc",
					Role:    "role",
				},
				Then: Then{
					Outcome: "allow",
				},
			},
		},
	}
	out := cfg.SerializeProtobuf()

	expected := &gconfigv1alpha1.Config{
		Admins: []*gconfigv1alpha1.Member{
			{
				Email: "admin@example.com",
			},
		},
		Users: []*gconfigv1alpha1.Member{
			{
				Email: "user@example.com",
			},
		},
		Groups: []*gconfigv1alpha1.Group{
			{
				Name: "group",
				Id:   "group",
				Members: []*gconfigv1alpha1.Member{
					{
						Email: "user@example.com",
					},
				},
			},
		},
		Providers: []*gconfigv1alpha1.Provider{
			{
				Id:               "provider",
				Type:             "awsRole",
				BastionAccountId: accID,
			},
		},
		Accounts: []*gconfigv1alpha1.Account{
			{
				Id:       "accgroup",
				Name:     "account group",
				Provider: "provider",
				Children: []*gconfigv1alpha1.Account{
					{
						Id:           "acc",
						Name:         "account",
						AwsAccountId: accID,
					},
				},
			},
		},
		Roles: []*gconfigv1alpha1.Role{
			{
				Id:       "role",
				Accounts: []string{"acc"},
				Policy:   "policy",
				Rules: []*gconfigv1alpha1.Rule{
					{
						Policy:          "allow",
						Group:           "test",
						SessionDuration: durationpb.New(time.Hour),
					},
				},
			},
		},
		Tests: []*gconfigv1alpha1.Test{
			{
				Name: "test",
				Given: &gconfigv1alpha1.Given{
					User:    "test@example.com",
					Account: "acc",
					Role:    "role",
				},
				Then: &gconfigv1alpha1.Then{
					Outcome: "allow",
				},
			},
		},
	}
	assert.Equal(t, expected, out)
}
