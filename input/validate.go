package input

// ValidateTokenInputBody represents the input to authenticate a JWT access token.
type ValidateTokenInputBody struct {
	AccessToken string `json:"access_token" doc:"The access token to authenticate"`
}

// NewValidateTokenInputBody creates a pointer to an input.ValidateTokenInputBody
func NewValidateTokenInputBody(accessToken string) *ValidateTokenInputBody {
	return &ValidateTokenInputBody{AccessToken: accessToken}
}
