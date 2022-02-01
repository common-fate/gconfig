package gconfig

import (
	"crypto/x509"
	"sort"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
)

type RulePolicy int

// The order of this enumeration should represent the order of preference for a developer
// e.g when assuming a role, which has multiple rules, the rule with assume will be selected before requireReason etc
//go:generate go run github.com/alvaroloes/enumer -type=RulePolicy -linecomment
const (
	RulePolicyAllow           RulePolicy = iota + 1 // allow
	RulePolicyRequireReason                         // requireReason
	RulePolicyRequireApproval                       // requireApproval
)

// This function is to be used by all services where there is a need to select a rule to apply from a list of rules
// For added security, this function takes in the user certificate to ensure that the groups match the rules
// The certificate should be validated before being used with this method
func RuleSelector(cert *x509.Certificate, rulesInput []*gconfigv1alpha1.Rule) *gconfigv1alpha1.Rule {
	admin := false
	groups := cert.Subject.OrganizationalUnit
	for _, group := range groups {
		if group == "granted:administrators" {
			admin = true
		}
	}
	rules := []*gconfigv1alpha1.Rule{}
	if admin {
		rules = rulesInput
	} else {
		for _, rule := range rulesInput {
			for _, group := range groups {
				if group == rule.Group {
					rules = append(rules, rule)
				}
			}
		}
	}

	if len(rules) == 0 {
		return nil
	}
	sort.Slice(rules, func(i, j int) bool {
		// error should never happen here
		a, err := RulePolicyString(rules[i].Policy)
		if err != nil {
			return false
		}
		// error should never happen here
		b, err := RulePolicyString(rules[j].Policy)
		if err != nil {
			return false
		}
		// uses teh ordering of the ENUM to determin order of rules
		return a < b
	})
	return rules[0]
}
