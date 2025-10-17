package public

import (
	"git.championtourney.com/championtourney/public/action"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OneTimeCodeTestSuite struct {
	suite.Suite
	expectedAction   string
	expectedMetaData map[string]any
}

func (testSuite *OneTimeCodeTestSuite) SetupTest() {
	testSuite.expectedAction = "CREATE_TOURNAMENT"
	testSuite.expectedMetaData = map[string]any{"some": "data"}
}

func (testSuite *OneTimeCodeTestSuite) TestNewOneTimeCodeMetadata_Success() {
	oneTimeCodeMetadata := NewOneTimeCodeMetadata(testSuite.expectedAction, testSuite.expectedMetaData)

	testSuite.Equal(testSuite.expectedAction, oneTimeCodeMetadata.Action)
	testSuite.Equal(testSuite.expectedMetaData, oneTimeCodeMetadata.MetaData)
}

func (testSuite *OneTimeCodeTestSuite) TestGetAction_Success() {
	expectedAction, err := action.MakeAction(testSuite.expectedAction)
	testSuite.Nil(err)

	oneTimeCodeMetadata := NewOneTimeCodeMetadata(testSuite.expectedAction, testSuite.expectedMetaData)

	theAction, err := oneTimeCodeMetadata.GetAction()
	testSuite.Nil(err)
	testSuite.Equal(expectedAction, theAction)
}

func Test_RunOneTimeCodeTestSuite(t *testing.T) {
	suite.Run(t, new(OneTimeCodeTestSuite))
}
