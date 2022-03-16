package gconfigdocs

import (
	"testing"

	gcoktav1alpha1 "github.com/common-fate/gconfig/pkg/okta/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func Test_PolicyParses(t *testing.T) {
	config, err := gcoktav1alpha1.ParseFile("reference_config.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, config)
}
