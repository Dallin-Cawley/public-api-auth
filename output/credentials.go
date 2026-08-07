package output

// CreateCredentialsOutputBody provides the structure of the output body in the POST /credentials request
type CreateCredentialsOutputBody struct {
	ClientID     string `json:"client_id" doc:"The generated client ID"`
	ClientSecret string `json:"client_secret" doc:"The generated client secret"`
}

// NewCreateCredentialsOutputBody creates the output body for the POST /credentials request
func NewCreateCredentialsOutputBody(clientID, clientSecret string) *CreateCredentialsOutputBody {
	return &CreateCredentialsOutputBody{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}
