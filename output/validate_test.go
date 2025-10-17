package output

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthenticateTestSuite struct {
	suite.Suite
	expectedApplicationID int
}

func (testSuite *AuthenticateTestSuite) SetupTest() {
	testSuite.expectedApplicationID = 1
}

func (testSuite *AuthenticateTestSuite) TestNewAuthenticateTokenOutputBody_Success() {
	outputBody := NewValidateTokenOutputBody(testSuite.expectedApplicationID)

	testSuite.Equal(testSuite.expectedApplicationID, outputBody.ApplicationID)
}

func Test_RunAuthenticateTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticateTestSuite))
}
