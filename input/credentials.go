package input

import (
	"fmt"

	"github.com/Dallin-Cawley/public-api-auth/auth"
	"github.com/Dallin-Cawley/public-api-auth/grant"
)

// CreateCredentialsInputBody represents the information necessary to generate a new client_id/client_secret
// pair for use with either the Client Credentials or Authorization Code OAuth2.0 flows.
type CreateCredentialsInputBody struct {
	RedirectURIs            []string       `json:"redirect_uris" schema:"redirect_uris" doc:"Array of redirection URIs"`
	TokenEndpointAuthMethod auth.Method    `json:"token_endpoint_auth_method" schema:"token_endpoint_auth_method" doc:"Requested authentication method for the token endpoint"`
	GrantTypes              grant.Types    `json:"grant_types" schema:"grant_types" enum:"client_credentials,authorization_code" doc:"The types of grant the credentials will be used for"`
	ResponseTypes           []string       `json:"response_types" schema:"response_types" doc:"Array of OAuth 2.0 response type strings"`
	ClientName              string         `json:"client_name" schema:"client_name" doc:"Human-readable name of the client"`
	ClientURI               string         `json:"client_uri" schema:"client_uri" doc:"URL of a web page providing information about the client"`
	LogoURI                 string         `json:"logo_uri" schema:"logo_uri" doc:"URL that references a logo for the client application"`
	Scope                   string         `json:"scope" schema:"scope" doc:"String containing a space-separated list of scope values"`
	Scopes                  []string       `json:"scopes" schema:"scopes" doc:"Legacy array of scope values"`
	Contacts                []string       `json:"contacts" schema:"contacts" doc:"Array of strings, each containing an email address"`
	TosURI                  string         `json:"tos_uri" schema:"tos_uri" doc:"URL that references a copy of the client's terms of service"`
	PolicyURI               string         `json:"policy_uri" schema:"policy_uri" doc:"URL that references a copy of the client's privacy policy"`
	JwksURI                 string         `json:"jwks_uri" schema:"jwks_uri" doc:"URL for the client's JSON Web Key Set [JWK] document"`
	Jwks                    map[string]any `json:"jwks" schema:"jwks" doc:"Client's JSON Web Key Set [JWK] document, passed by value"`
	SoftwareID              string         `json:"software_id" schema:"software_id" doc:"A unique identifier string assigned by the client developer"`
	SoftwareVersion         string         `json:"software_version" schema:"software_version" doc:"A version identifier string for the client software"`
}

// CreateCredentialsInputOption represents a functional option for configuring a CreateCredentialsInputBody
type CreateCredentialsInputOption func(*CreateCredentialsInputBody)

// WithAuthorizationCodeGrantType adds the authorization_code grant type to the credentials, and requires
// a list of redirection URIs.
func WithAuthorizationCodeGrantType(redirectURIs []string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.GrantTypes = append(body.GrantTypes, grant.AuthorizationCode)
		body.RedirectURIs = redirectURIs
	}
}

// WithClientCredentialsGrantType adds the client_credentials grant type to the credentials.
func WithClientCredentialsGrantType() CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.GrantTypes = append(body.GrantTypes, grant.ClientCredentials)
	}
}

// WithTokenEndpointAuthMethod sets the token_endpoint_auth_method on the CreateCredentialsInputBody.
func WithTokenEndpointAuthMethod(method auth.Method) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.TokenEndpointAuthMethod = method
	}
}

// WithResponseTypes sets the response_types on the CreateCredentialsInputBody.
func WithResponseTypes(responseTypes []string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.ResponseTypes = responseTypes
	}
}

// WithClientName sets the client_name on the CreateCredentialsInputBody.
func WithClientName(name string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.ClientName = name
	}
}

// WithClientURI sets the client_uri on the CreateCredentialsInputBody.
func WithClientURI(uri string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.ClientURI = uri
	}
}

// WithLogoURI sets the logo_uri on the CreateCredentialsInputBody.
func WithLogoURI(uri string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.LogoURI = uri
	}
}

// WithScope sets the scope string on the CreateCredentialsInputBody.
func WithScope(scope string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.Scope = scope
	}
}

// WithScopes sets the legacy scopes array on the CreateCredentialsInputBody.
func WithScopes(scopes []string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.Scopes = scopes
	}
}

// WithContacts sets the contacts on the CreateCredentialsInputBody.
func WithContacts(contacts []string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.Contacts = contacts
	}
}

// WithTosURI sets the tos_uri on the CreateCredentialsInputBody.
func WithTosURI(uri string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.TosURI = uri
	}
}

// WithPolicyURI sets the policy_uri on the CreateCredentialsInputBody.
func WithPolicyURI(uri string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.PolicyURI = uri
	}
}

// WithJwksURI sets the jwks_uri on the CreateCredentialsInputBody.
func WithJwksURI(uri string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.JwksURI = uri
	}
}

// WithJwks sets the jwks on the CreateCredentialsInputBody.
func WithJwks(jwks map[string]any) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.Jwks = jwks
	}
}

// WithSoftwareID sets the software_id on the CreateCredentialsInputBody.
func WithSoftwareID(id string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.SoftwareID = id
	}
}

// WithSoftwareVersion sets the software_version on the CreateCredentialsInputBody.
func WithSoftwareVersion(version string) CreateCredentialsInputOption {
	return func(body *CreateCredentialsInputBody) {
		body.SoftwareVersion = version
	}
}

// NewCreateCredentialsInputBody creates a pointer to an input.CreateCredentialsInputBody
func NewCreateCredentialsInputBody(opts ...CreateCredentialsInputOption) *CreateCredentialsInputBody {
	body := &CreateCredentialsInputBody{}

	for _, opt := range opts {
		opt(body)
	}

	if body.TokenEndpointAuthMethod == auth.MethodUnknown {
		body.TokenEndpointAuthMethod = auth.ClientSecretBasic
	}

	return body
}

// GetGrantTypes retrieves the requested grant types.
func (body *CreateCredentialsInputBody) GetGrantTypes() (grant.Types, error) {
	for _, gt := range body.GrantTypes {
		if gt == grant.TypeUnknown {
			return nil, fmt.Errorf("invalid grant type")
		}
	}
	return body.GrantTypes, nil
}

// GetTokenEndpointAuthMethod retrieves the requested token endpoint authentication method.
func (body *CreateCredentialsInputBody) GetTokenEndpointAuthMethod() (auth.Method, error) {
	if body.TokenEndpointAuthMethod == auth.MethodUnknown {
		return auth.MethodUnknown, fmt.Errorf("invalid token endpoint auth method")
	}
	return body.TokenEndpointAuthMethod, nil
}
