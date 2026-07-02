package input

import "github.com/Dallin-Cawley/public-api-auth/grant"

// ClientCredentialsPayload represents a client credentials grant payload.
type ClientCredentialsPayload struct {
	GrantType grant.Type `json:"grant_type"`
	OwnerID   string     `json:"owner_id"`
}

// NewClientCredentialsPayload creates a new instance of ClientCredentialsPayload with the provided grantType and ownerID.
func NewClientCredentialsPayload(ownerID string) *ClientCredentialsPayload {
	return &ClientCredentialsPayload{
		GrantType: grant.ClientCredentials,
		OwnerID:   ownerID,
	}
}
