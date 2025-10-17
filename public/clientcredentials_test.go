package public

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientCredentialsTestSuite struct {
	suite.Suite
	expectedClientID     string
	expectedClientSecret string
}

func (testSuite *ClientCredentialsTestSuite) SetupTest() {
	testSuite.expectedClientID = "some id"
	testSuite.expectedClientSecret = "some secret"
}

func (testSuite *ClientCredentialsTestSuite) TestNewClientCredentialsGrant_Success() {
	clientCredentials := NewClientCredentialsGrant(testSuite.expectedClientID, testSuite.expectedClientSecret)

	testSuite.Equal(testSuite.expectedClientID, clientCredentials.ClientID)
	testSuite.Equal(testSuite.expectedClientSecret, clientCredentials.ClientSecret)
}

func Test_RunClientCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(ClientCredentialsTestSuite))
}
