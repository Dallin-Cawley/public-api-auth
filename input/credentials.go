package input

import (
	"github.com/Dallin-Cawley/public-api-auth/grant"
)

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
func (body *CreateCredentialsInputBody) GetGrantTypes() (grant.Types, error) {
	return grant.TypesFromString(body.GrantTypes)
}
