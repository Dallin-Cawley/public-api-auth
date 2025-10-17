package output

import (
	"git.championtourney.com/championtourney/public-api-auth/public"
)

// ClientCredentialsOutputBody represents the response body for Oauth2.0 client credentials grant.
type ClientCredentialsOutputBody struct {
	Grant public.ClientCredentialsGrant `json:"grant" doc:"The client credentials grant for the application"`
}

// NewClientCredentialsOutputBody creates a new instance of outputClientCredentialsOutputBody with the provided public.ClientCredentialsGrant.
func NewClientCredentialsOutputBody(grant public.ClientCredentialsGrant) *ClientCredentialsOutputBody {
	return &ClientCredentialsOutputBody{
		Grant: grant,
	}
}
