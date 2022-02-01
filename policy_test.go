package gconfig

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"testing"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"github.com/stretchr/testify/assert"
)

// Tests the the correct ordering of rules is observed
func Test_RuleSelector(t *testing.T) {
	cert := &x509.Certificate{Subject: pkix.Name{OrganizationalUnit: []string{"developer", "tester"}}}
	rules := []*gconfigv1alpha1.Rule{{Policy: RulePolicyRequireReason.String(), Group: "developer"}, {Policy: RulePolicyAllow.String(), Group: "tester"}, {Policy: RulePolicyRequireApproval.String(), Group: "reasonNeeders"}}
	rule, _ := RuleSelector(cert, rules)
	assert.Equal(t, rules[1], rule)

	// user does not have tester group so the developer rule is retured
	cert = &x509.Certificate{Subject: pkix.Name{OrganizationalUnit: []string{"developer"}}}
	rule, _ = RuleSelector(cert, rules)
	assert.Equal(t, rules[0], rule)
}
