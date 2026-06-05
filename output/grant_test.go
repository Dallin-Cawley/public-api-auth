package output

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GrantTestSuite struct {
	suite.Suite
}

func (testSuite *GrantTestSuite) Test_NewClientCredentialsOutputBody_Success() {
	clientID := "some id"
	clientSecret := "some secret"
	body := NewClientCredentialsGrant(clientID, clientSecret)

	testSuite.NotNil(body)
	testSuite.Equal(clientID, body.ClientID)
	testSuite.Equal(clientSecret, body.ClientSecret)
}

func Test_RunGrantTestSuite(t *testing.T) {
	suite.Run(t, new(GrantTestSuite))
}
