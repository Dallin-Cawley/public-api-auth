package input

import (
	"testing"

	"git.championtourney.com/championtourney/public-api-auth/grant"

	"github.com/stretchr/testify/suite"
)

type GrantTestSuite struct {
	suite.Suite
}

func (testSuite *GrantTestSuite) Test_NewClientCredentialsPayload_Success() {
	ownerID := "test-owner-id"
	payload := NewClientCredentialsPayload(ownerID)

	testSuite.NotNil(payload)
	testSuite.Equal(ownerID, payload.OwnerID)
	testSuite.Equal(grant.ClientCredentials, payload.GrantType)
}

func Test_RunGrantTestSuite(t *testing.T) {
	suite.Run(t, new(GrantTestSuite))
}
