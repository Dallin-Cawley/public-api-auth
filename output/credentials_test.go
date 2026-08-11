package output

import (
	"encoding/json"
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

func (testSuite *CredentialsTestSuite) TestCreateCredentialsOutputBody_JSONMarshaling_OmitsEmptyFields() {
	outputBody := &CreateCredentialsOutputBody{
		ClientID:     "some client id",
		ClientSecret: "some client secret",
	}

	jsonData, err := json.Marshal(outputBody)
	testSuite.NoError(err)

	var result map[string]any
	err = json.Unmarshal(jsonData, &result)
	testSuite.NoError(err)

	testSuite.Equal("some client id", result["client_id"])
	testSuite.Equal("some client secret", result["client_secret"])

	// Fields that should be missing
	missingFields := []string{
		"redirect_uris",
		"token_endpoint_auth_method",
		"grant_types",
		"response_types",
		"client_name",
		"client_uri",
		"logo_uri",
		"scopes",
		"contacts",
		"tos_uri",
		"policy_uri",
		"jwks_uri",
		"jwks",
		"software_id",
		"software_version",
	}

	for _, field := range missingFields {
		_, exists := result[field]
		testSuite.False(exists, "Field %s should be missing from JSON", field)
	}
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
