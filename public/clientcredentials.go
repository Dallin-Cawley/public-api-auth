package public

type ClientCredentialsGrant struct {
	ClientID     string `json:"client_id" doc:"The client ID of the application"`
	ClientSecret string `json:"client_secret" doc:"The client secret of the application"`
}

// NewClientCredentialsGrant creates a public.ClientCredentialsGrant
func NewClientCredentialsGrant(clientID string, clientSecret string) *ClientCredentialsGrant {
	return &ClientCredentialsGrant{ClientID: clientID, ClientSecret: clientSecret}
}
