package output

type ClientCredentialsGrant struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// NewClientCredentialsGrant creates a new output.ClientCredentialsGrant
func NewClientCredentialsGrant(clientID, clientSecret string) *ClientCredentialsGrant {
	return &ClientCredentialsGrant{ClientID: clientID, ClientSecret: clientSecret}
}
