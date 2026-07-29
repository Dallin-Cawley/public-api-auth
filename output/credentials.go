package output

// CreateCredentialsOutputBody provides the structure of the output body in the POST /credentials request
type CreateCredentialsOutputBody struct {
	ClientID     string   `json:"client_id" doc:"The generated client ID"`
	ClientSecret string   `json:"client_secret" doc:"The generated client secret"`
	GrantTypes   []string `json:"grant_types" doc:"The type of grant the credentials were created for"`
	Scopes       []string `json:"scopes" doc:"The scopes that can be attached to tokens generated from these credentials"`
}

// NewCreateCredentialsOutputBody creates the output body for the POST /credentials request
func NewCreateCredentialsOutputBody(clientID, clientSecret string, grantTypes, scopes []string) *CreateCredentialsOutputBody {
	return &CreateCredentialsOutputBody{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantTypes:   grantTypes,
		Scopes:       scopes,
	}
}
