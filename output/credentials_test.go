package output

import (
	"testing"

	"github.com/Dallin-Cawley/public-api-auth/auth"
	"github.com/Dallin-Cawley/public-api-auth/grant"
	"github.com/Dallin-Cawley/public-api-auth/response"
	"github.com/stretchr/testify/suite"
)

type CredentialsTestSuite struct {
	suite.Suite
}

func (testSuite *CredentialsTestSuite) TestNewCreateCredentialsOutputBody_Success() {
	expectedClientID := "some client id"
	expectedClientSecret := "some client secret"
	expectedRedirectURIs := []string{"http://localhost:8080/callback"}
	expectedTokenEndpointAuthMethod := auth.ClientSecretBasic
	expectedGrantTypes := grant.Types{grant.ClientCredentials}
	expectedResponseTypes := response.Types{response.Code}
	expectedClientName := "some client name"
	expectedClientURI := "http://localhost:8080"
	expectedLogoURI := "http://localhost:8080/logo.png"
	expectedScopes := []string{"read", "write"}
	expectedContacts := []string{"test@example.com"}
	expectedTosURI := "http://localhost:8080/tos"
	expectedPolicyURI := "http://localhost:8080/policy"
	expectedJwksURI := "http://localhost:8080/jwks"
	expectedJwks := map[string]any{"key": "value"}
	expectedSoftwareID := "some software id"
	expectedSoftwareVersion := "some software version"

	outputBody := NewCreateCredentialsOutputBody(
		expectedClientID,
		expectedClientSecret,
		expectedRedirectURIs,
		expectedTokenEndpointAuthMethod,
		expectedGrantTypes,
		expectedResponseTypes,
		expectedClientName,
		expectedClientURI,
		expectedLogoURI,
		expectedScopes,
		expectedContacts,
		expectedTosURI,
		expectedPolicyURI,
		expectedJwksURI,
		expectedJwks,
		expectedSoftwareID,
		expectedSoftwareVersion,
	)

	testSuite.Equal(expectedClientID, outputBody.ClientID)
	testSuite.Equal(expectedClientSecret, outputBody.ClientSecret)
	testSuite.Equal(expectedRedirectURIs, outputBody.RedirectURIs)
	testSuite.Equal(expectedTokenEndpointAuthMethod, outputBody.TokenEndpointAuthMethod)
	testSuite.Equal(expectedGrantTypes, outputBody.GrantTypes)
	testSuite.Equal(expectedResponseTypes, outputBody.ResponseTypes)
	testSuite.Equal(expectedClientName, outputBody.ClientName)
	testSuite.Equal(expectedClientURI, outputBody.ClientURI)
	testSuite.Equal(expectedLogoURI, outputBody.LogoURI)
	testSuite.Equal(expectedScopes, outputBody.Scopes)
	testSuite.Equal(expectedContacts, outputBody.Contacts)
	testSuite.Equal(expectedTosURI, outputBody.TosURI)
	testSuite.Equal(expectedPolicyURI, outputBody.PolicyURI)
	testSuite.Equal(expectedJwksURI, outputBody.JwksURI)
	testSuite.Equal(expectedJwks, outputBody.Jwks)
	testSuite.Equal(expectedSoftwareID, outputBody.SoftwareID)
	testSuite.Equal(expectedSoftwareVersion, outputBody.SoftwareVersion)
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
