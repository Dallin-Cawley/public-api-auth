package output

import (
	"github.com/Dallin-Cawley/public-api-auth/auth"
	"github.com/Dallin-Cawley/public-api-auth/grant"
	"github.com/Dallin-Cawley/public-api-auth/response"
)

// CreateCredentialsOutputBody provides the structure of the output body in the POST /credentials request
type CreateCredentialsOutputBody struct {
	ClientID                string         `json:"client_id" doc:"The generated client ID"`
	ClientSecret            string         `json:"client_secret" doc:"The generated client secret"`
	RedirectURIs            []string       `json:"redirect_uris,omitempty" schema:"redirect_uris" doc:"Array of redirection URIs"`
	TokenEndpointAuthMethod auth.Method    `json:"token_endpoint_auth_method" schema:"token_endpoint_auth_method" doc:"Requested authentication method for the token endpoint"`
	GrantTypes              grant.Types    `json:"grant_types" schema:"grant_types" enum:"client_credentials,authorization_code" doc:"The types of grant the credentials will be used for"`
	ResponseTypes           response.Types `json:"response_types" schema:"response_types" enum:"code,token" doc:"The types of response the credentials will be used for"`
	ClientName              string         `json:"client_name" schema:"client_name" doc:"Human-readable name of the client"`
	ClientURI               string         `json:"client_uri" schema:"client_uri" doc:"URL of a web page providing information about the client"`
	LogoURI                 string         `json:"logo_uri" schema:"logo_uri" doc:"URL that references a logo for the client application"`
	Scopes                  []string       `json:"scopes" schema:"scopes" doc:"list of scope values"`
	Contacts                []string       `json:"contacts" schema:"contacts" doc:"Array of strings, each containing an email address"`
	TosURI                  string         `json:"tos_uri" schema:"tos_uri" doc:"URL that references a copy of the client's terms of service"`
	PolicyURI               string         `json:"policy_uri" schema:"policy_uri" doc:"URL that references a copy of the client's privacy policy"`
	JwksURI                 string         `json:"jwks_uri" schema:"jwks_uri" doc:"URL for the client's JSON Web Key Set [JWK] document"`
	Jwks                    map[string]any `json:"jwks" schema:"jwks" doc:"Client's JSON Web Key Set [JWK] document, passed by value"`
	SoftwareID              string         `json:"software_id" schema:"software_id" doc:"A unique identifier string assigned by the client developer"`
	SoftwareVersion         string         `json:"software_version" schema:"software_version" doc:"A version identifier string for the client software"`
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
		Scopes:                  scopes,
		Contacts:                contacts,
		TosURI:                  tosURI,
		PolicyURI:               policyURI,
		JwksURI:                 jwksURI,
		Jwks:                    jwks,
		SoftwareID:              softwareID,
		SoftwareVersion:         softwareVersion,
	}
}
