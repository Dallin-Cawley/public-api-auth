package input

import (
	"testing"

	"git.championtourney.com/championtourney/public/action"

	"github.com/stretchr/testify/suite"
)

type OneTimeCodeTestSuite struct {
	suite.Suite
}

func (testSuite *OneTimeCodeTestSuite) TestNewEmptyCreateOneTimeCodeInputBody_Success() {
	body := NewEmptyCreateOneTimeCodeInputBody()

	testSuite.Empty(body.Action)
	testSuite.Empty(body.MetaData)
}

func (testSuite *OneTimeCodeTestSuite) TestNewCreateOneTimeCodeInputBody_Success() {
	expectedAction := action.CreateTeam
	expectedMetaData := map[string]string{"some": "data"}

	body := NewCreateOneTimeCodeInputBody(expectedAction, expectedMetaData)

	testSuite.Equal(expectedAction.String(), body.Action)
	testSuite.Equal(expectedMetaData, body.MetaData)
}

func (testSuite *OneTimeCodeTestSuite) TestGetAction_Success() {
	expectedAction := action.CreateTeam

	body := NewCreateOneTimeCodeInputBody(expectedAction, nil)
	theAction, err := body.GetAction()

	testSuite.Nil(err)
	testSuite.Equal(expectedAction, theAction)
}

func Test_RunOneTimeCodeTestSuite(t *testing.T) {
	suite.Run(t, new(OneTimeCodeTestSuite))
}
