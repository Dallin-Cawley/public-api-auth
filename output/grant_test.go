package output

import (
	"git.championtourney.com/championtourney/public-api-auth/public"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GrantTestSuite struct {
	suite.Suite
	clientCredentials *public.ClientCredentialsGrant
}

func (testSuite *GrantTestSuite) SetupTest() {
	testSuite.clientCredentials = public.NewClientCredentialsGrant("some id", "some secret")
}

func (testSuite *GrantTestSuite) TestNewClientCredentialsOutputBody_Success() {
	body := NewClientCredentialsOutputBody(*testSuite.clientCredentials)

	testSuite.Equal(*testSuite.clientCredentials, body.Grant)
}

func Test_RunGrantTestSuite(t *testing.T) {
	suite.Run(t, new(GrantTestSuite))
}
