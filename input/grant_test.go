package input

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type GrantTestSuite struct {
	suite.Suite
	expectedApplicationName string
	expectedCallbackURL     string
}

func (testSuite *GrantTestSuite) SetupTest() {
	testSuite.expectedApplicationName = "some name"
	testSuite.expectedCallbackURL = "some url"
}

func (testSuite *GrantTestSuite) TestNewEmptyClientCredentialsGrantBody_Success() {
	body := NewEmptyClientCredentialsInputBody()

	testSuite.Empty(body.ApplicationName)
	testSuite.Nil(body.CallbackURL)
}

func (testSuite *GrantTestSuite) TestNewClientCredentialsGrantBody_Success() {
	body := NewClientCredentialsInputBody(testSuite.expectedApplicationName, testSuite.expectedCallbackURL)

	testSuite.Equal(testSuite.expectedApplicationName, body.ApplicationName)
	testSuite.Equal(testSuite.expectedCallbackURL, *body.CallbackURL)
}

func Test_RunGrantTestSuite(t *testing.T) {
	suite.Run(t, new(GrantTestSuite))
}
