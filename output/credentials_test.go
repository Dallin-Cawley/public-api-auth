package output

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CredentialsTestSuite struct {
	suite.Suite
}

func (testSuite *CredentialsTestSuite) TestNewCreateCredentialsOutputBody_Success() {
	expectedClientID := "some client id"
	expectedClientSecret := "some client secret"
	expectedGrantType := []string{"client_credentials"}
	expectedScopes := []string{"some scope"}

	outputBody := NewCreateCredentialsOutputBody(expectedClientID, expectedClientSecret, expectedGrantType, expectedScopes)

	testSuite.Equal(expectedClientID, outputBody.ClientID)
	testSuite.Equal(expectedClientSecret, outputBody.ClientSecret)
	testSuite.Equal(expectedGrantType, outputBody.GrantTypes)
	testSuite.Equal(expectedScopes, outputBody.Scopes)
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
