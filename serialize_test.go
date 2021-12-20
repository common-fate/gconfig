package gconfig

import (
	"testing"
	"time"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	cfg = Config{
		providers: &gconfigv1alpha1.Providers{
			Providers: []*gconfigv1alpha1.Provider{
				{
					Id: "aws",
					Accounts: []*gconfigv1alpha1.Account{
						{
							Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
							Id:   "acc",
						},
					},
				},
			},
		},
		Type: "granted/v1alpha1",
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
		Roles: []*Role{
			{
				ID:       "role",
				Accounts: []string{"acc"},
				roleAccounts: []RoleAccount{
					{
						AccountID:  "acc",
						ProviderID: "aws",
					},
				},
				Policy: "policy",
				Rules: []Rule{
					{
						Policy:          "allow",
						Group:           "test",
						SessionDuration: time.Hour,
					},
				},
			},
		},
		Tests: []Test{
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

	expected = &gconfigv1alpha1.Config{
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
		Roles: []*gconfigv1alpha1.Role{
			{
				Id: "role",
				Accounts: []*gconfigv1alpha1.RoleAccount{
					{
						ProviderId: "aws",
						AccountId:  "acc",
					},
				},
				Policy: "policy",
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
)

func TestSerialize(t *testing.T) {
	out := cfg.SerializeProtobuf()
	assert.Equal(t, expected, out)
}

func TestDeserialize(t *testing.T) {
	reversed := FromProtobuf(expected, cfg.providers)
	assert.Equal(t, cfg, reversed)
}
