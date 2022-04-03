package gconfig

import (
	"testing"
	"time"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	cfg = Config{
		providers: &gconfigv1alpha1.Providers{
			Providers: []*gconfigv1alpha1.Provider{
				{
					Id: "aws",
					Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
						{
							Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
							Id:   "acc",
						},
					}}},
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
				Type:     "aws",
				Accounts: []Account{{Account: "acc"}},
				roleAccounts: []RoleAccount{
					{
						AccountID:  "acc",
						ProviderID: "aws",
					},
				},
				SessionDuration: time.Hour,
				Policy:          "policy",
				Rules: []Rule{
					{
						Policy: map[string]interface{}{
							"allow": true,
						},
						Group: "test",
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
				Id:   "role",
				Type: "aws",
				Accounts: []*gconfigv1alpha1.RoleAccount{
					{
						ProviderId: "aws",
						AccountId:  "acc",
					},
				},
				SessionDuration: durationpb.New(time.Hour),
				Policy:          "policy",
				Rules: []*gconfigv1alpha1.Rule{
					{
						Policy: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"allow": structpb.NewBoolValue(true),
							},
						},
						Group: "test",
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
	out, err := cfg.SerializeProtobuf()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, out)
}

func TestDeserialize(t *testing.T) {
	reversed := FromProtobuf(expected, cfg.providers)
	assert.Equal(t, cfg, reversed)
}
