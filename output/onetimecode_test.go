package output

import (
	"testing"

	"github.com/Dallin-Cawley/public-api-auth/public"

	"github.com/stretchr/testify/suite"
)

type OnetimeCodeTestSuite struct {
	suite.Suite
	metaData *public.OneTimeCodeMetadata
}

func (testSuite *OnetimeCodeTestSuite) SetupTest() {
	testSuite.metaData = public.NewOneTimeCodeMetadata("CREATE_TOURNAMENT", map[string]any{"some": "data"})
}

func (testSuite *OnetimeCodeTestSuite) TestNewGetOneTimeCodeMetaDataOutputBody_Success() {
	outputBody := NewGetOneTimeCodeMetaDataOutputBody(testSuite.metaData)

	testSuite.Equal(*testSuite.metaData, outputBody.MetaData)
}

func (testSuite *OnetimeCodeTestSuite) TestNewCreateOneTimeCodeOutputBody_Success() {
	expectedOneTimeCode := "some code"

	outputBody := NewCreateOneTimeCodeOutputBody(expectedOneTimeCode)

	testSuite.Equal(expectedOneTimeCode, outputBody.OneTimeCode)
}

func Test_RunOneTimeCodeTestSuite(t *testing.T) {
	suite.Run(t, new(OnetimeCodeTestSuite))
}
