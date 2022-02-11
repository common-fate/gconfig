package gconfig

import (
	"bytes"
	"fmt"
	"reflect"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"gopkg.in/yaml.v3"
)

// providerWrapperYAML is a wrapper around providers which allows us to recursively parse and combine
// providers together in different files, into a single YAML file.
type providerWrapperYAML struct {
	Version   string                  `yaml:"version"`
	Providers map[string]providerYAML `yaml:"providers"`
}

// providerYAML is the YAML representation of
// a provider config.
type providerYAML struct {
	Type                string              `yaml:"type"`
	Name                string              `yaml:"name"`
	ManagementAccountID *string             `yaml:"managementAccountId,omitempty"`
	AccessHandlers      []accessHandlerYAML `yaml:"accessHandlers"`
	Accounts            []accountYAML       `yaml:"accounts"`
}

type accountYAML struct {
	ID       string        `yaml:"id"`
	Name     *string       `yaml:"name,omitempty"`
	Type     string        `yaml:"type"`
	Accounts []accountYAML `yaml:"accounts,omitempty"`
}

type accessHandlerYAML struct {
	URL string `yaml:"url"`
}

// ProvidersToYAML serializes the Provider object to a YAML string
func ProvidersToYAML(p *gconfigv1alpha1.Provider) ([]byte, error) {
	pyaml := providerYAML{
		Name: p.Name,
	}

	switch v := p.Details.(type) {
	case *gconfigv1alpha1.Provider_Aws:
		pyaml.ManagementAccountID = &v.Aws.OrgManagementAccountId
		pyaml.Type = "aws"
	case *gconfigv1alpha1.Provider_AwsSso:
		pyaml.ManagementAccountID = &v.AwsSso.OrgManagementAccountId
		pyaml.Type = "awsSSO"
	default:
		return nil, fmt.Errorf("unhandled provider type %s", reflect.TypeOf(p.Details))
	}

	for _, acc := range p.Accounts {
		ayaml := buildAccountYAML(acc)
		pyaml.Accounts = append(pyaml.Accounts, ayaml)
	}

	for _, ah := range p.AccessHandlers {
		handler := accessHandlerYAML{
			URL: ah.Url,
		}
		pyaml.AccessHandlers = append(pyaml.AccessHandlers, handler)
	}

	wrapper := providerWrapperYAML{
		Version: "granted/providers/v1alpha1",
		Providers: map[string]providerYAML{
			p.Id: pyaml,
		},
	}

	var out bytes.Buffer

	fmt.Fprintf(&out, "# This file has been automatically generated by Granted. DO NOT EDIT.\n")
	fmt.Fprintf(&out, "# To enrol accounts belonging to this provider with Granted, run \"granted accounts enrol --provider %s\" (https://granted.dev/adding-accounts)\n", p.Id)

	enc := yaml.NewEncoder(&out)
	enc.SetIndent(2)
	err := enc.Encode(wrapper)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func buildAccountYAML(a *gconfigv1alpha1.Account) accountYAML {
	ayaml := accountYAML{
		ID: a.Id,
	}

	switch a.Type {
	case gconfigv1alpha1.Account_TYPE_AWS_ACCOUNT:
		ayaml.Type = "AWS::Account"
	}

	if a.Name != "" {
		ayaml.Name = &a.Name
	}

	for _, child := range a.Children {
		childyaml := buildAccountYAML(child)
		ayaml.Accounts = append(ayaml.Accounts, childyaml)
	}

	return ayaml
}
