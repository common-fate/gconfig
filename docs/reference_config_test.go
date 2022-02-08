package gconfigdocs

import (
	"testing"

	"github.com/common-fate/gconfig"
	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func Test_PolicyParses(t *testing.T) {
	config, err := gconfig.ParseFile("reference_config.yaml", &gconfigv1alpha1.Providers{Providers: []*gconfigv1alpha1.Provider{
		{
			Id: "ac1",
			Accounts: []*gconfigv1alpha1.Account{
				{
					Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
					Id:      "123456789120",
					Aliases: []string{"sandbox"},
				},
			},
		},
		{
			Id: "ac2",
			Accounts: []*gconfigv1alpha1.Account{
				{
					Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
					Id:   "123456789121",
				},
			},
		},
		{
			Id: "ac3",
			Accounts: []*gconfigv1alpha1.Account{
				{
					Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
					Id:      "123456789122",
					Aliases: []string{"sandbox"},
				},
			},
		},
		{
			Id: "ac4",
			Accounts: []*gconfigv1alpha1.Account{
				{
					Type:    gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
					Id:      "123456789123",
					Aliases: []string{"research"},
				},
			},
		},
	}})
	assert.NoError(t, err)
	assert.NotNil(t, config)
}
