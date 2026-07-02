package input

import (
	"testing"

	"github.com/Dallin-Cawley/public-api-auth/grant"
	"github.com/stretchr/testify/suite"
)

type TokenTestSuite struct {
	suite.Suite
}

func (testSuite *TokenTestSuite) TestNewEmptyCreateClientCredentialsTokenInputBody_Success() {
	body := NewEmptyCreateTokenInputBody()

	testSuite.Empty(body)
}

func (testSuite *TokenTestSuite) TestNewCreateClientCredentialsTokenInputBody_Success() {
	expectedClientID := "some id"
	expectedClientSecret := "some secret"

	body := NewCreateClientCredentialsTokenInputBody(&expectedClientID, &expectedClientSecret)

	testSuite.Equal(&expectedClientID, body.ClientID)
	testSuite.Equal(&expectedClientSecret, body.ClientSecret)
	testSuite.Equal(grant.ClientCredentials.String(), body.GrantType)
}

func (testSuite *TokenTestSuite) TestGetGrantType_Success() {
	expectedGrantType := grant.ClientCredentials

	body := NewCreateClientCredentialsTokenInputBody(nil, nil)
	theGrantType, err := body.GetGrantType()

	testSuite.Nil(err)
	testSuite.Equal(expectedGrantType, theGrantType)
}

func Test_RunTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
