package input

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ScopeTestSuite struct {
	suite.Suite
}

func (testSuite *ScopeTestSuite) Test_NewCreateScopesInputBody() {
	expectedScopes := []string{"scope1", "scope2"}
	body := NewCreateScopesInputBody(expectedScopes)

	testSuite.Equal(expectedScopes, body.Scopes)
}

func Test_RunScopeTestSuite(t *testing.T) {
	suite.Run(t, new(ScopeTestSuite))
}
