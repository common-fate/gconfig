package gconfig

import (
	"io/ioutil"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"gopkg.in/yaml.v3"
)

// ParseFile parses a Granted YAML config file and implements read time checks
// See validate.go for any checks on the config post reading the file
func ParseFile(filename string, providers *gconfigv1alpha1.Providers) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContents(filename, b, providers)
}

// parseContents parses a Granted YAML config file
func parseContents(filename string, in []byte, providers *gconfigv1alpha1.Providers) (*Config, error) {
	var c Config

	c.providers = providers

	err := yaml.Unmarshal(in, &c)
	if err != nil {
		return nil, err
	}

	// add the filename to entities in the config
	for _, a := range c.Admins {
		a.pos.Filename = filename
	}

	for _, u := range c.Users {
		u.pos.Filename = filename
	}

	for _, g := range c.Groups {
		g.pos.Filename = filename

		for _, m := range g.Members {
			m.pos.Filename = filename
		}
	}

	for _, r := range c.Roles {
		r.pos.Filename = filename

		// this is all provider-specific validation, so skip it for now.
		// roleType := r.Type
		// isOktaRole := roleType == gconfigv1alpha1.RoleType_ROLE_TYPE_OKTA
		// if isOktaRole {
		// 	if r.Group == "" {
		// 		err = fmt.Errorf("group required on each role")
		// 		err = printLintError(r, err)
		// 		return nil, err
		// 	}
		// 	if r.Policy != "" {
		// 		err = fmt.Errorf("policy not supported on Okta roles")
		// 		err = printLintError(r, err)
		// 		return nil, err
		// 	}
		// 	if len(r.Accounts) != 0 {
		// 		err = fmt.Errorf("accounts not supported on Okta roles")
		// 		err = printLintError(r, err)
		// 		return nil, err
		// 	}
		// 	if r.SessionDuration != 0 {
		// 		err = fmt.Errorf("session duration not supported on Okta roles")
		// 		err = printLintError(r, err)
		// 		return nil, err
		// 	}
		// 	if r.DefaultRegion != "" {
		// 		err = fmt.Errorf("default region not supported on Okta roles")
		// 		err = printLintError(r, err)
		// 		return nil, err
		// 	}
		// } else {
		// 	if r.SessionDuration <= 0 {
		// 		err = fmt.Errorf("session required on each role")
		// 		err = printLintError(r, err)
		// 		return nil, err
		// 	}
		// 	if r.Group != "" {
		// 		err = fmt.Errorf("group not supported on AWS roles")
		// 		err = printLintError(r, err)
		// 		return nil, err
		// 	}
		// }

		for _, rule := range r.Rules {
			rule.Policy.pos.Filename = filename
			// // Validates that the rule policy matches a supported policy type
			// if policy, err := RulePolicyString(rule.Policy.Policy); err != nil {
			// 	policyValues := []string{}
			// 	for _, pol := range RulePolicyValues() {
			// 		policyValues = append(policyValues, pol.String())
			// 	}
			// 	err = fmt.Errorf("policy: %s must be one of %v", rule.Policy.Policy, policyValues)
			// 	err = printLintError(&rule.Policy, err)
			// 	return nil, err

			// 	//the breakglass field is allowed on roles with requireApproval and allows users to bypass the approval step (but sends an alert to Granted that they have done so)
			// } else if rule.Breakglass && policy != RulePolicyRequireApproval {
			// 	err = fmt.Errorf("'breakglass: true' can only be used on policies which require approval")
			// 	err = printLintError(&rule.Policy, err)
			// 	return nil, err
			// }
		}
	}

	err = c.setRoleAccounts()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
