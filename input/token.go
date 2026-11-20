package input

import (
	"strings"

	"github.com/Dallin-Cawley/public-api-auth/grant"
)

// CreateTokenInputBody is a token request (RFC 6749 Section 4.1.3 and 4.4.2): a client_credentials grant, or the
// exchange of an authorization code (authorization_code grant with PKCE, RFC 7636 Section 4.5).
type CreateTokenInputBody struct {
	ClientID     string `json:"client_id" schema:"client_id" doc:"The client ID, when the client doesn't authenticate with HTTP Basic"`
	ClientSecret string `json:"client_secret" schema:"client_secret" doc:"The client secret, for client_secret_post"`
	GrantType    string `json:"grant_type" schema:"grant_type" enum:"client_credentials,authorization_code,refresh_token" doc:"The type of grant to issue"`
	Scope        string `json:"scope" schema:"scope" doc:"The scope of the token as a space delimited list of scopes (client_credentials; refresh_token, where it may only narrow the granted scope)"`
	Code         string `json:"code" schema:"code" doc:"The authorization code the authorization endpoint returned (authorization_code)"`
	CodeVerifier string `json:"code_verifier" schema:"code_verifier" doc:"The PKCE code verifier of the authorization request (authorization_code)"`
	RedirectURI  string `json:"redirect_uri" schema:"redirect_uri" doc:"The redirect_uri of the authorization request, when it included one (authorization_code)"`
	RefreshToken string `json:"refresh_token" schema:"refresh_token" doc:"The refresh token to exchange (refresh_token, RFC 6749 Section 6)"`
}

// NewCreateTokenInputBody creates a pointer to an input.CreateTokenInputBody
func NewCreateTokenInputBody(clientID, clientSecret string) *CreateTokenInputBody {
	return &CreateTokenInputBody{ClientID: clientID, ClientSecret: clientSecret, GrantType: grant.ClientCredentials.String()}
}

// GetGrantType converts the input.CreateTokenInputBody to its grant.Type counterpart.
func (body *CreateTokenInputBody) GetGrantType() (grant.Type, error) {
	return grant.MakeType(body.GrantType)
}

// Scopes returns the scopes of the token as a slice of strings.
func (body *CreateTokenInputBody) Scopes() []string {
	return strings.Split(body.Scope, " ")
}
