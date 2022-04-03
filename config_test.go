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

// func Test_PolicyValidated(t *testing.T) {

// 	b, err := ioutil.ReadFile("configtest.yaml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rule1 := `    rules:
//     - policy: badPolicyName
//       group: developers
//      `
// 	cf := b
// 	_, err = parseContents("configtest.yaml", append(cf, []byte(rule1)...), &gconfigv1alpha1.Providers{})
// 	policyValues := []string{}
// 	for _, pol := range RulePolicyValues() {
// 		policyValues = append(policyValues, pol.String())
// 	}
// 	expected := fmt.Sprintf("configtest.yaml:32:15: policy: badPolicyName must be one of %v", policyValues)
// 	assert.EqualError(t, err, expected)

// 	rule2 := `    rules:
//     - policy: requireApproval
//       group: developers
//       `

// 	cf = b
// 	providers := &gconfigv1alpha1.Providers{Providers: []*gconfigv1alpha1.Provider{
// 		{
// 			Id: "aws",
// 			Details: &gconfigv1alpha1.Provider_Aws{Aws: &gconfigv1alpha1.AWSProviderDetails{Accounts: []*gconfigv1alpha1.Account{
// 				{
// 					Type: gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT,
// 					Id:   "123456789012",
// 				},
// 			}}},
// 		},
// 	}}
// 	c, err := parseContents("configtest.yaml", append(cf, []byte(rule2)...), providers)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, c)
// 	assert.Equal(t, c.Roles[0].Rules[0].Policy.Policy, RulePolicyRequireApproval.String())

// 	rule3 := `    rules:
//     - policy: allow
//       breakglass: true
//       group: developers
//       sessionDuration: 8h`
// 	cf = b
// 	_, err = parseContents("configtest.yaml", append(cf, []byte(rule3)...), providers)

// 	expected = "configtest.yaml:32:15: 'breakglass: true' can only be used on policies which require approval"
// 	assert.EqualError(t, err, expected)

// 	rule4 := `    rules:
//     - policy: requireApproval
//       breakglass: true
//       group: developers
//       sessionDuration: 8h`

// 	cf = b

// 	c, err = parseContents("configtest.yaml", append(cf, []byte(rule4)...), providers)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, c)
// 	assert.Equal(t, c.Roles[0].Rules[0].Policy.Policy, RulePolicyRequireApproval.String())
// }
