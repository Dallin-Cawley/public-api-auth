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

	outputBody := NewCreateTokenOutputBody(expectedAccessToken, expectedAccessExpires)

	testSuite.Equal(expectedAccessToken, outputBody.AccessToken)
	testSuite.Equal(expectedAccessExpires, outputBody.AccessExpires)
}

func Test_RunTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
