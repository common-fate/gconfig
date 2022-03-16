package gcoktav1alpha1

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// ParseFile parses a Granted YAML config file and implements read time checks
// See validate.go for any checks on the config post reading the file
func ParseFile(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return parseContents(filename, b)
}

// parseContents parses a Granted YAML config file
func parseContents(filename string, in []byte) (*Config, error) {
	var c Config

	err := yaml.Unmarshal(in, &c)
	if err != nil {
		return nil, err
	}

	// add the filename to entities in the config
	for _, a := range c.Admins {
		a.pos.Filename = filename
	}

	for _, r := range c.Roles {
		r.pos.Filename = filename
		for _, rule := range r.Rules {
			rule.Policy.pos.Filename = filename
			// Validates that the rule policy matches a supported policy type
			if policy, err := RulePolicyString(rule.Policy.Policy); err != nil {
				policyValues := []string{}
				for _, pol := range RulePolicyValues() {
					policyValues = append(policyValues, pol.String())
				}
				err = fmt.Errorf("policy: %s must be one of %v", rule.Policy.Policy, policyValues)
				err = printLintError(&rule.Policy, err)
				return nil, err

				//the breakglass field is allowed on roles with requireApproval and allows users to bypass the approval step (but sends an alert to Granted that they have done so)
			} else if rule.Breakglass && policy != RulePolicyRequireApproval {
				err = fmt.Errorf("'breakglass: true' can only be used on policies which require approval")
				err = printLintError(&rule.Policy, err)
				return nil, err
			}
		}

	}

	return &c, nil
}

// filePositioners can get their positions in a YAML file.
// all of our config nodes are filePositioners
type filePositioner interface {
	filePosition() *FilePosition
}

// printLintError adds file position information to the error message if it exists.
func printLintError(p filePositioner, err error) error {
	pos := p.filePosition()
	if pos == nil {
		return err
	}
	return fmt.Errorf("%s:%d:%d: %s", pos.Filename, pos.Line, pos.Col, err)
}
