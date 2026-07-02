package input

import (
	"testing"

	"github.com/Dallin-Cawley/public-api-auth/grant"
	"github.com/stretchr/testify/suite"
)

type TokenTestSuite struct {
	suite.Suite
}

func (testSuite *TokenTestSuite) TestNewCreateTokenInputBody_Success() {
	expectedClientID := "some id"
	expectedClientSecret := "some secret"

	body := NewCreateTokenInputBody(&expectedClientID, &expectedClientSecret)

	testSuite.Equal(&expectedClientID, body.ClientID)
	testSuite.Equal(&expectedClientSecret, body.ClientSecret)
	testSuite.Equal(grant.ClientCredentials.String(), body.GrantType)
}

func (testSuite *TokenTestSuite) TestGetGrantType_Success() {
	expectedGrantType := grant.ClientCredentials

	body := NewCreateTokenInputBody(nil, nil)
	theGrantType, err := body.GetGrantType()

	testSuite.NoError(err)
	testSuite.Equal(expectedGrantType, theGrantType)
}

func (testSuite *TokenTestSuite) TestGetGrantType_Failure() {
	body := &CreateTokenInputBody{GrantType: "invalid"}
	_, err := body.GetGrantType()

	testSuite.ErrorContains(err, "invalid grant type")
}

func Test_RunTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
