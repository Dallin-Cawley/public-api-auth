package input

import (
	"testing"

	"github.com/Dallin-Cawley/public-api-auth/grant"
	"github.com/stretchr/testify/suite"
)

type CredentialsTestSuite struct {
	suite.Suite
}

func (testSuite *CredentialsTestSuite) TestNewCreateCredentialsInputBody_Success() {
	body := NewCreateCredentialsInputBody(grant.AuthorizationCode)

	testSuite.Equal([]string{grant.AuthorizationCode.String()}, body.GrantType)
}

func (testSuite *CredentialsTestSuite) TestGetGrantType_Success() {
	body := NewCreateCredentialsInputBody(grant.ClientCredentials)
	theGrantType, err := body.GetGrantType()

	testSuite.NoError(err)
	testSuite.Equal([]grant.Type{grant.ClientCredentials}, theGrantType)
}

func (testSuite *CredentialsTestSuite) TestGetGrantType_Failure() {
	body := &CreateCredentialsInputBody{GrantType: []string{"invalid"}}
	_, err := body.GetGrantType()

	testSuite.ErrorContains(err, "invalid grant type")
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
