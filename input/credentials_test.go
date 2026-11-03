package input

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

func (testSuite *CredentialsTestSuite) TestNewCreateCredentialsInputBody_Success() {
	body := NewCreateCredentialsInputBody(
		WithAuthorizationCodeGrantType("https://example.com/callback"),
		WithScopes("one"),
	)

	testSuite.Equal(grant.Types{grant.AuthorizationCode}, body.GrantTypes)
	testSuite.Equal([]string{"https://example.com/callback"}, body.RedirectURIs)
	testSuite.Equal("one", body.Scope)
	testSuite.Equal([]string{"one"}, body.Scopes())
}

func (testSuite *CredentialsTestSuite) TestNewCreateCredentialsInputBody_BothGrants() {
	body := NewCreateCredentialsInputBody(
		WithAuthorizationCodeGrantType("https://example.com/callback"),
		WithClientCredentialsGrantType(),
	)

	testSuite.Contains(body.GrantTypes, grant.AuthorizationCode)
	testSuite.Contains(body.GrantTypes, grant.ClientCredentials)
	testSuite.Equal([]string{"https://example.com/callback"}, body.RedirectURIs)
}

func (testSuite *CredentialsTestSuite) TestCreateCredentialsInputBody_GetGrantTypes_Success() {
	body := NewCreateCredentialsInputBody(
		WithClientCredentialsGrantType(),
		WithScopes("one"),
	)
	theGrantType := body.GetGrantTypes()

	testSuite.Equal(grant.Types{grant.ClientCredentials}, theGrantType)
}

func (testSuite *CredentialsTestSuite) TestAllOptions() {
	jwks := map[string]any{"key": "value"}
	body := NewCreateCredentialsInputBody(
		WithTokenEndpointAuthMethod(auth.ClientSecretPost),
		WithResponseTypes(response.Code),
		WithClientName("my-client"),
		WithClientURI("https://example.com"),
		WithLogoURI("https://example.com/logo.png"),
		WithScopes("openid", "profile"),
		WithContacts("admin@example.com"),
		WithTosURI("https://example.com/tos"),
		WithPolicyURI("https://example.com/policy"),
		WithJwksURI("https://example.com/jwks"),
		WithJwks(jwks),
		WithSoftwareID("soft-123"),
		WithSoftwareVersion("1.0.0"),
	)

	testSuite.Equal(auth.ClientSecretPost, body.TokenEndpointAuthMethod)
	testSuite.Equal(response.Types{response.Code}, body.ResponseTypes)
	testSuite.Equal("my-client", body.ClientName)
	testSuite.Equal("https://example.com", body.ClientURI)
	testSuite.Equal("https://example.com/logo.png", body.LogoURI)
	testSuite.Equal("openid profile", body.Scope)
	testSuite.Equal([]string{"openid", "profile"}, body.Scopes())
	testSuite.Equal([]string{"admin@example.com"}, body.Contacts)
	testSuite.Equal("https://example.com/tos", body.TosURI)
	testSuite.Equal("https://example.com/policy", body.PolicyURI)
	testSuite.Equal("https://example.com/jwks", body.JwksURI)
	testSuite.Equal(jwks, body.Jwks)
	testSuite.Equal("soft-123", body.SoftwareID)
	testSuite.Equal("1.0.0", body.SoftwareVersion)
}

func (testSuite *CredentialsTestSuite) TestCreateCredentialsInputBody_GetResponseTypes_Success() {
	body := NewCreateCredentialsInputBody(
		WithResponseTypes(response.Code),
	)
	theResponseType := body.GetResponseTypes()

	testSuite.Equal(response.Types{response.Code}, theResponseType)
}

func (testSuite *CredentialsTestSuite) TestGetTokenEndpointAuthMethod_Success() {
	body := NewCreateCredentialsInputBody(
		WithTokenEndpointAuthMethod(auth.ClientSecretPost),
	)
	theMethod := body.GetTokenEndpointAuthMethod()

	testSuite.Equal(auth.ClientSecretPost, theMethod)
}

func (testSuite *CredentialsTestSuite) TestCreateCredentialsInputBody_Marshaling() {
	body := NewCreateCredentialsInputBody(
		WithTokenEndpointAuthMethod(auth.ClientSecretPost),
		WithAuthorizationCodeGrantType("https://example.com/callback"),
		WithResponseTypes(response.Code),
	)

	// Marshal
	data, err := json.Marshal(body)
	testSuite.NoError(err)

	// Assert JSON contains expected values
	jsonStr := string(data)
	testSuite.Contains(jsonStr, `"token_endpoint_auth_method":"client_secret_post"`)
	testSuite.Contains(jsonStr, `"grant_types":["authorization_code"]`)
	testSuite.Contains(jsonStr, `"response_types":["code"]`)

	// Unmarshal
	var unmarshaled CreateCredentialsInputBody
	err = json.Unmarshal(data, &unmarshaled)
	testSuite.NoError(err)

	// Assert unmarshaled values
	testSuite.Equal(body.TokenEndpointAuthMethod, unmarshaled.TokenEndpointAuthMethod)
	testSuite.Equal(body.GrantTypes, unmarshaled.GrantTypes)
	testSuite.Equal(body.ResponseTypes, unmarshaled.ResponseTypes)
	testSuite.Equal(body.RedirectURIs, unmarshaled.RedirectURIs)
}

func (testSuite *CredentialsTestSuite) TestScopes_SplitsOnWhitespace() {
	var body CreateCredentialsInputBody
	testSuite.NoError(json.Unmarshal([]byte(`{"scope":"  read\twrite  admin "}`), &body))
	testSuite.Equal([]string{"read", "write", "admin"}, body.Scopes())

	testSuite.Empty((&CreateCredentialsInputBody{}).Scopes())
}

func (testSuite *CredentialsTestSuite) TestUnmarshalJSON_CollectsLocalized() {
	var body CreateCredentialsInputBody
	err := json.Unmarshal([]byte(`{"client_name":"n","client_name#ja-Jpan-JP":"名前","client_uri#fr":"https://x/fr","scope":"a"}`), &body)
	testSuite.Require().NoError(err)

	testSuite.Equal("n", body.ClientName)
	testSuite.Equal("a", body.Scope)
	testSuite.Equal(map[string]string{"client_name#ja-Jpan-JP": "名前", "client_uri#fr": "https://x/fr"}, map[string]string(body.Localized))
}

func (testSuite *CredentialsTestSuite) TestUnmarshalJSON_TypeErrorsKeepTheField() {
	var body CreateCredentialsInputBody

	var typeErr *json.UnmarshalTypeError
	testSuite.Require().ErrorAs(json.Unmarshal([]byte(`{"contacts":"not-a-list"}`), &body), &typeErr)
	testSuite.Equal("contacts", typeErr.Field)

	testSuite.Require().ErrorAs(json.Unmarshal([]byte(`{"client_name#fr":1}`), &body), &typeErr)
	testSuite.Equal("client_name#fr", typeErr.Field)
}

func (testSuite *CredentialsTestSuite) TestMarshalJSON_RoundTripsLocalized() {
	body := NewCreateCredentialsInputBody(WithClientCredentialsGrantType(), WithClientName("n"), WithLocalized("client_name", "fr", "Nom"))

	data, err := json.Marshal(body)
	testSuite.Require().NoError(err)

	var members map[string]any
	testSuite.Require().NoError(json.Unmarshal(data, &members))
	testSuite.Equal("Nom", members["client_name#fr"])
	testSuite.NotContains(members, "Localized")

	var decoded CreateCredentialsInputBody
	testSuite.Require().NoError(json.Unmarshal(data, &decoded))
	testSuite.Equal(body.Localized, decoded.Localized)
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
