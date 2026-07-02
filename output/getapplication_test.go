package output

import (
	"testing"

	"github.com/Dallin-Cawley/public-api-auth/public"
	"github.com/stretchr/testify/suite"
)

type GetApplicationTestSuite struct {
	suite.Suite
}

func (testSuite *GetApplicationTestSuite) TestNewGetApplicationOutputBody_Success() {
	publicApplication := public.NewApplication(1, "some name", "some id", "some secret", "some callback url", nil)

	outputBody := NewGetApplicationOutputBody(*publicApplication)
	testSuite.Equal(*publicApplication, outputBody.Application)
}

func Test_RunGetApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(GetApplicationTestSuite))
}
