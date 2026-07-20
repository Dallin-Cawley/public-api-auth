package input

import (
	"fmt"

	"github.com/Dallin-Cawley/public-api-auth/grant"
)

// CreateCredentialsInputBody represents the information necessary to generate a new client_id/client_secret
// pair for use with either the Client Credentials or Authorization Code OAuth2.0 flows.
type CreateCredentialsInputBody struct {
	GrantType []string `json:"grant_type" schema:"grant_type" enum:"client_credentials,authorization_code" doc:"The type of grant the credentials will be used for"`
}

// NewCreateCredentialsInputBody creates a pointer to an input.CreateCredentialsInputBody
func NewCreateCredentialsInputBody(grantTypes ...grant.Type) *CreateCredentialsInputBody {
	grantTypeStrings := make([]string, len(grantTypes))
	for i, grantType := range grantTypes {
		grantTypeStrings[i] = grantType.String()
	}

	return &CreateCredentialsInputBody{GrantType: grantTypeStrings}
}

// GetGrantType converts the input.CreateCredentialsInputBody to its grant.Type counterpart.
func (body *CreateCredentialsInputBody) GetGrantType() ([]grant.Type, error) {
	grantTypes := make([]grant.Type, len(body.GrantType))
	for i, grantType := range body.GrantType {
		var err error
		grantTypes[i], err = grant.MakeType(grantType)
		if err != nil {
			return nil, fmt.Errorf("failed to convert grant type: %w", err)
		}
	}

	return grantTypes, nil
}
