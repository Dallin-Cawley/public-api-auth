package output

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidateTestSuite struct {
	suite.Suite
}

func (testSuite *ValidateTestSuite) TestNewValidateOutputBody_Success() {
	expectedSubject := "some-subject"
	expectedIssuedAt := "some-iat"
	expectedExpiresAt := "some-exp"
	expectedJWTID := "some-jti"
	expectedScopes := []string{"scope1", "scope2"}

	outputBody := NewValidateOutputBody(expectedSubject, expectedIssuedAt, expectedExpiresAt, expectedJWTID, expectedScopes)

	testSuite.Equal(expectedSubject, outputBody.Subject)
	testSuite.Equal(expectedIssuedAt, outputBody.IssuedAt)
	testSuite.Equal(expectedExpiresAt, outputBody.ExpiresAt)
	testSuite.Equal(expectedJWTID, outputBody.JWTID)
	testSuite.Equal(expectedScopes, outputBody.Scopes)
}

func Test_RunValidateTestSuite(t *testing.T) {
	suite.Run(t, new(ValidateTestSuite))
}
