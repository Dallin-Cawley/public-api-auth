package output

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TokenTestSuite struct {
	suite.Suite
}

func (testSuite *TokenTestSuite) TestNewCreateTokenOutputBody_Success() {
	expectedAccessToken := "some access token"
	expectedAccessExpires := "some access expires"
	expectedTokenType := "Bearer"

	outputBody := NewCreateTokenOutputBody(expectedAccessToken, expectedAccessExpires)

	testSuite.Equal(expectedAccessToken, outputBody.AccessToken)
	testSuite.Equal(expectedAccessExpires, outputBody.AccessExpires)
	testSuite.Equal(expectedTokenType, outputBody.TokenType)
}

func (testSuite *TokenTestSuite) TestWithRefreshTokenAndScope() {
	outputBody := NewCreateTokenOutputBody("a", "b").WithRefreshToken("r").WithScope("read")

	testSuite.Equal("r", outputBody.RefreshToken)
	testSuite.Equal("read", outputBody.Scope)
}

func Test_RunTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
