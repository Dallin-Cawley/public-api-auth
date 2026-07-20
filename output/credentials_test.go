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
	expectedGrantType := "client_credentials"

	outputBody := NewCreateCredentialsOutputBody(expectedClientID, expectedClientSecret, expectedGrantType)

	testSuite.Equal(expectedClientID, outputBody.ClientID)
	testSuite.Equal(expectedClientSecret, outputBody.ClientSecret)
	testSuite.Equal([]string{expectedGrantType}, outputBody.GrantTypes)
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
