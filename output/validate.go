package output

// ValidateTokenOutputBody contains the structure of the output body for the /oauth/token/validate endpoint
type ValidateTokenOutputBody struct {
	ApplicationID int `json:"application_id" doc:"The application ID for whom the token was generated"`
}

// NewValidateTokenOutputBody creates the output body for the /oauth/token/validate endpoint
func NewValidateTokenOutputBody(applicationID int) *ValidateTokenOutputBody {
	return &ValidateTokenOutputBody{ApplicationID: applicationID}
}
