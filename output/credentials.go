package output

import (
	"strings"

	"github.com/Dallin-Cawley/public-api-auth/auth"
	"github.com/Dallin-Cawley/public-api-auth/grant"
	"github.com/Dallin-Cawley/public-api-auth/response"
)

// CreateCredentialsOutputBody provides the structure of the output body in the POST /credentials request
type CreateCredentialsOutputBody struct {
	ClientID                string         `json:"client_id" doc:"The generated client ID"`
	ClientSecret            string         `json:"client_secret" doc:"The generated client secret"`
	ClientIDIssuedAt        int64          `json:"client_id_issued_at,omitempty" doc:"Time at which the client_id was issued, in seconds since the Unix epoch (RFC 7591 Section 3.2.1)"`
	ClientSecretExpiresAt   int64          `json:"client_secret_expires_at" doc:"Time at which the client_secret expires, in seconds since the Unix epoch, or 0 if it never expires (RFC 7591 Section 3.2.1)"`
	RegistrationAccessToken string         `json:"registration_access_token,omitempty" doc:"Token used to access the client configuration endpoint (RFC 7592 Section 3)"`
	RegistrationClientURI   string         `json:"registration_client_uri,omitempty" doc:"Fully qualified URL of the client configuration endpoint for this client (RFC 7592 Section 3)"`
	RedirectURIs            []string       `json:"redirect_uris,omitempty" schema:"redirect_uris" doc:"Array of redirection URIs"`
	TokenEndpointAuthMethod auth.Method    `json:"token_endpoint_auth_method,omitempty" schema:"token_endpoint_auth_method" doc:"Requested authentication method for the token endpoint"`
	GrantTypes              grant.Types    `json:"grant_types,omitempty" schema:"grant_types" enum:"client_credentials,authorization_code" doc:"The types of grant the credentials will be used for"`
	ResponseTypes           response.Types `json:"response_types,omitempty" schema:"response_types" enum:"code,token" doc:"The types of response the credentials will be used for"`
	ClientName              string         `json:"client_name,omitempty" schema:"client_name" doc:"Human-readable name of the client"`
	ClientURI               string         `json:"client_uri,omitempty" schema:"client_uri" doc:"URL of a web page providing information about the client"`
	LogoURI                 string         `json:"logo_uri,omitempty" schema:"logo_uri" doc:"URL that references a logo for the client application"`
	Scope                   string         `json:"scope,omitempty" schema:"scope" doc:"Space-separated list of scope values (RFC 7591 Section 2)"`
	Contacts                []string       `json:"contacts,omitempty" schema:"contacts" doc:"Array of strings, each containing an email address"`
	TosURI                  string         `json:"tos_uri,omitempty" schema:"tos_uri" doc:"URL that references a copy of the client's terms of service"`
	PolicyURI               string         `json:"policy_uri,omitempty" schema:"policy_uri" doc:"URL that references a copy of the client's privacy policy"`
	JwksURI                 string         `json:"jwks_uri,omitempty" schema:"jwks_uri" doc:"URL for the client's JSON Web Key Set [JWK] document"`
	Jwks                    map[string]any `json:"jwks,omitempty" schema:"jwks" doc:"Client's JSON Web Key Set [JWK] document, passed by value"`
	SoftwareID              string         `json:"software_id,omitempty" schema:"software_id" doc:"A unique identifier string assigned by the client developer"`
	SoftwareVersion         string         `json:"software_version,omitempty" schema:"software_version" doc:"A version identifier string for the client software"`
}

// NewCreateCredentialsOutputBody creates the output body for the POST /credentials request
func NewCreateCredentialsOutputBody(
	clientID, clientSecret string,
	redirectURIs []string,
	tokenEndpointAuthMethod auth.Method,
	grantTypes grant.Types,
	responseTypes response.Types,
	clientName, clientURI, logoURI string,
	scopes []string,
	contacts []string,
	tosURI, policyURI, jwksURI string,
	jwks map[string]any,
	softwareID string, softwareVersion string,
) *CreateCredentialsOutputBody {
	return &CreateCredentialsOutputBody{
		ClientID:                clientID,
		ClientSecret:            clientSecret,
		RedirectURIs:            redirectURIs,
		TokenEndpointAuthMethod: tokenEndpointAuthMethod,
		GrantTypes:              grantTypes,
		ResponseTypes:           responseTypes,
		ClientName:              clientName,
		ClientURI:               clientURI,
		LogoURI:                 logoURI,
		Scope:                   strings.Join(scopes, " "),
		Contacts:                contacts,
		TosURI:                  tosURI,
		PolicyURI:               policyURI,
		JwksURI:                 jwksURI,
		Jwks:                    jwks,
		SoftwareID:              softwareID,
		SoftwareVersion:         softwareVersion,
	}
}
