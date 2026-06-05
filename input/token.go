package input

import "git.championtourney.com/championtourney/public-api-auth/grant"

// CreateTokenInputBody represents the information necessary to create a token from information
// generated via the Client Credentials Oauth2.0 flow.
type CreateTokenInputBody struct {
	ClientID     *string `json:"client_id" schema:"client_id" doc:"The client ID of the application when using the Client Credentials flow"`
	ClientSecret *string `json:"client_secret" schema:"client_secret" doc:"The client secret of the application when using the Client Credentials flow"`
	GrantType    string  `json:"grant_type" schema:"grant_type" enum:"client_credentials" doc:"The type of grant to issue"`
}

// NewEmptyCreateTokenInputBody creates an empty input.CreateTokenInputBody
func NewEmptyCreateTokenInputBody() *CreateTokenInputBody {
	return &CreateTokenInputBody{}
}

// NewCreateClientCredentialsTokenInputBody creates an input.CreateTokenInputBody with the necessary information
// to create a JWT token.
func NewCreateClientCredentialsTokenInputBody(clientID, clientSecret *string) *CreateTokenInputBody {
	return &CreateTokenInputBody{ClientID: clientID, ClientSecret: clientSecret, GrantType: grant.ClientCredentials.String()}
}

// GetGrantType converts the input.CreateTokenInputBody to its grant.Type counterpart.
func (body *CreateTokenInputBody) GetGrantType() (grant.Type, error) {
	return grant.MakeType(body.GrantType)
}
