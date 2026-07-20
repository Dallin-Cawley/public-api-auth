package input

import (
	"github.com/Dallin-Cawley/public-api-auth/grant"
)

// InternalCreateCredentialsInputBody is used to create a new ClientCredentials in api-auth for internal use. It's useful
// when you want to create human-readable credentials in a staging environment.
type InternalCreateCredentialsInputBody struct {
	ClientID     string   `json:"client_id" schema:"client_id" doc:"The client_id of the credentials"`
	ClientSecret string   `json:"client_secret" schema:"client_secret" doc:"The client_secret of the credentials"`
	GrantTypes   []string `json:"grant_types" schema:"grant_types" enum:"client_credentials,authorization_code" doc:"The types of grant the credentials will be used for"`
}

// NewInternalCreateCredentialsInputBody creates a pointer to an input.InternalCreateCredentialsInputBody
func NewInternalCreateCredentialsInputBody(grantTypes grant.Types, clientID, clientSecret string) *InternalCreateCredentialsInputBody {
	return &InternalCreateCredentialsInputBody{GrantTypes: grantTypes.Strings(), ClientID: clientID, ClientSecret: clientSecret}
}

// GetGrantTypes retrieves the requested grant types.
func (body *InternalCreateCredentialsInputBody) GetGrantTypes() ([]grant.Type, error) {
	return grant.TypesFromString(body.GrantTypes)
}

// CreateCredentialsInputBody represents the information necessary to generate a new client_id/client_secret
// pair for use with either the Client Credentials or Authorization Code OAuth2.0 flows.
type CreateCredentialsInputBody struct {
	GrantTypes []string `json:"grant_types" schema:"grant_types" enum:"client_credentials,authorization_code" doc:"The types of grant the credentials will be used for"`
}

// NewCreateCredentialsInputBody creates a pointer to an input.CreateCredentialsInputBody
func NewCreateCredentialsInputBody(grantTypes grant.Types) *CreateCredentialsInputBody {
	return &CreateCredentialsInputBody{GrantTypes: grantTypes.Strings()}
}

// GetGrantTypes retrieves the requested grant types.
func (body *CreateCredentialsInputBody) GetGrantTypes() ([]grant.Type, error) {
	return grant.TypesFromString(body.GrantTypes)
}
