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
	body := NewCreateCredentialsInputBody(grant.Types{grant.AuthorizationCode}, []string{"one"})

	testSuite.Equal([]string{grant.AuthorizationCode.String()}, body.GrantTypes)
	testSuite.Equal([]string{"one"}, body.Scopes)
}

func (testSuite *CredentialsTestSuite) TestCreateCredentialsInputBody_GetGrantTypes_Success() {
	body := NewCreateCredentialsInputBody(grant.Types{grant.ClientCredentials}, []string{"one"})
	theGrantType, err := body.GetGrantTypes()

	testSuite.NoError(err)
	testSuite.Equal(grant.Types{grant.ClientCredentials}, theGrantType)
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
