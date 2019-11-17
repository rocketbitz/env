package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnvTestSuite struct {
	suite.Suite
}

func (s *EnvTestSuite) TestEnv() {
	SetDefaults(map[string]string{
		"DEFAULT_KEY_ONE": "default_value_one",
		"DEFAULT_KEY_TWO": "default_value_two",
	})

	assert.Equal(s.T(), 2, DefaultCount())
	assert.GreaterOrEqual(s.T(), Count(), 2)
	assert.Equal(s.T(), Count()-DefaultCount(), ExplicitCount())

	SetDefault("DEFAULT_KEY_THREE", "default_value_three")
	assert.Equal(s.T(), "default_value_three", Get("DEFAULT_KEY_THREE"))

	Set("DEFAULT_KEY_ONE", "explicit_value_one")
	assert.Equal(s.T(), "explicit_value_one", Get("DEFAULT_KEY_ONE"))

	assert.Empty(s.T(), Get("UNSET_KEY"))

	Set("EXPLICIT_KEY_ONE", "explicit_value_one")
	assert.Equal(s.T(), "explicit_value_one", Get("EXPLICIT_KEY_ONE"))

	assert.Contains(s.T(), String(), `"key": "EXPLICIT_KEY_ONE`)
	assert.Contains(s.T(), String(), `"value": "explicit_value_one"`)

	assert.NotPanics(s.T(), Print)
}

func TestEnvTestSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}
