package input

import (
	"testing"

	"github.com/Dallin-Cawley/public-api-auth/grant"
	"github.com/stretchr/testify/suite"
)

type CredentialsTestSuite struct {
	suite.Suite
}

func (testSuite *CredentialsTestSuite) TestNewInternalCreateCredentialsInputBody_Success() {
	expectedClientID := "client_id"
	expectedClientSecret := "client_secret"

	body := NewInternalCreateCredentialsInputBody(grant.Types{grant.AuthorizationCode}, expectedClientID, expectedClientSecret)

	testSuite.Equal([]string{grant.AuthorizationCode.String()}, body.GrantTypes)
	testSuite.Equal(expectedClientID, body.ClientID)
	testSuite.Equal(expectedClientSecret, body.ClientSecret)
}

func (testSuite *CredentialsTestSuite) TestCreateInternalCreateCredentialsInputBody_GetGrantTypes_Success() {
	body := NewInternalCreateCredentialsInputBody(grant.Types{grant.AuthorizationCode}, "client_id", "client_secret")
	theGrantType, err := body.GetGrantTypes()

	testSuite.NoError(err)
	testSuite.Equal([]grant.Type{grant.AuthorizationCode}, theGrantType)
}

func (testSuite *CredentialsTestSuite) TestInternalCreateCredentialsInputBody_GetGrantTypes_Success() {
	body := &InternalCreateCredentialsInputBody{GrantTypes: []string{grant.ClientCredentials.String()}}
	theGrantType, err := body.GetGrantTypes()

	testSuite.NoError(err)
	testSuite.Equal([]grant.Type{grant.ClientCredentials}, theGrantType)
}

func (testSuite *CredentialsTestSuite) TestNewCreateCredentialsInputBody_Success() {
	body := NewCreateCredentialsInputBody(grant.Types{grant.AuthorizationCode})

	testSuite.Equal([]string{grant.AuthorizationCode.String()}, body.GrantTypes)
}

func (testSuite *CredentialsTestSuite) TestCreateCredentialsInputBody_GetGrantTypes_Success() {
	body := NewCreateCredentialsInputBody(grant.Types{grant.ClientCredentials})
	theGrantType, err := body.GetGrantTypes()

	testSuite.NoError(err)
	testSuite.Equal([]grant.Type{grant.ClientCredentials}, theGrantType)
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
