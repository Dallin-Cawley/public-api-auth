package input

import (
	"testing"

	"github.com/Dallin-Cawley/public-api-auth/grant"
	"github.com/stretchr/testify/suite"
)

type CredentialsTestSuite struct {
	suite.Suite
}

func (testSuite *CredentialsTestSuite) TestNewCreateCredentialsInputBody_Success() {
	body := NewCreateCredentialsInputBody(
		WithAuthorizationCodeGrantType([]string{"https://example.com/callback"}),
		WithScopes([]string{"one"}),
	)

	testSuite.Equal([]string{grant.AuthorizationCode.String()}, body.GrantTypes)
	testSuite.Equal([]string{"https://example.com/callback"}, body.RedirectURIs)
	testSuite.Equal([]string{"one"}, body.Scopes)
}

func (testSuite *CredentialsTestSuite) TestNewCreateCredentialsInputBody_BothGrants() {
	body := NewCreateCredentialsInputBody(
		WithAuthorizationCodeGrantType([]string{"https://example.com/callback"}),
		WithClientCredentialsGrantType(),
	)

	testSuite.Contains(body.GrantTypes, grant.AuthorizationCode.String())
	testSuite.Contains(body.GrantTypes, grant.ClientCredentials.String())
	testSuite.Equal([]string{"https://example.com/callback"}, body.RedirectURIs)
}

func (testSuite *CredentialsTestSuite) TestCreateCredentialsInputBody_GetGrantTypes_Success() {
	body := NewCreateCredentialsInputBody(
		WithClientCredentialsGrantType(),
		WithScopes([]string{"one"}),
	)
	theGrantType, err := body.GetGrantTypes()

	testSuite.NoError(err)
	testSuite.Equal(grant.Types{grant.ClientCredentials}, theGrantType)
}

func (testSuite *CredentialsTestSuite) TestAllOptions() {
	jwks := map[string]any{"key": "value"}
	body := NewCreateCredentialsInputBody(
		WithTokenEndpointAuthMethod("client_secret_post"),
		WithResponseTypes([]string{"code"}),
		WithClientName("my-client"),
		WithClientURI("https://example.com"),
		WithLogoURI("https://example.com/logo.png"),
		WithScope("openid profile"),
		WithScopes([]string{"openid", "profile"}),
		WithContacts([]string{"admin@example.com"}),
		WithTosURI("https://example.com/tos"),
		WithPolicyURI("https://example.com/policy"),
		WithJwksURI("https://example.com/jwks"),
		WithJwks(jwks),
		WithSoftwareID("soft-123"),
		WithSoftwareVersion("1.0.0"),
	)

	testSuite.Equal("client_secret_post", body.TokenEndpointAuthMethod)
	testSuite.Equal([]string{"code"}, body.ResponseTypes)
	testSuite.Equal("my-client", body.ClientName)
	testSuite.Equal("https://example.com", body.ClientURI)
	testSuite.Equal("https://example.com/logo.png", body.LogoURI)
	testSuite.Equal("openid profile", body.Scope)
	testSuite.Equal([]string{"openid", "profile"}, body.Scopes)
	testSuite.Equal([]string{"admin@example.com"}, body.Contacts)
	testSuite.Equal("https://example.com/tos", body.TosURI)
	testSuite.Equal("https://example.com/policy", body.PolicyURI)
	testSuite.Equal("https://example.com/jwks", body.JwksURI)
	testSuite.Equal(jwks, body.Jwks)
	testSuite.Equal("soft-123", body.SoftwareID)
	testSuite.Equal("1.0.0", body.SoftwareVersion)
}

func (testSuite *CredentialsTestSuite) TestGetGrantTypes_Failure() {
	body := &CreateCredentialsInputBody{
		GrantTypes: []string{"invalid_grant"},
	}
	_, err := body.GetGrantTypes()
	testSuite.Error(err)
}

func Test_RunCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
