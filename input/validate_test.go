package input

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthenticateTestSuite struct {
	suite.Suite
	expectedToken string
}

func (testSuite *AuthenticateTestSuite) SetupTest() {
	testSuite.expectedToken = "some token"
}

func (testSuite *AuthenticateTestSuite) TestNewValidateTokenInputBody_Empty() {
	body := NewValidateTokenInputBody("")

	testSuite.Empty(body.AccessToken)
}

func (testSuite *AuthenticateTestSuite) TestNewAuthenticateTokenBody_Success() {
	body := NewValidateTokenInputBody(testSuite.expectedToken)

	testSuite.Equal(testSuite.expectedToken, body.AccessToken)
}

func Test_RunAuthenticateTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticateTestSuite))
}
