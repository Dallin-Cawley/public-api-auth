package input

// ValidateTokenInputBody represents the input to authenticate a JWT access token.
type ValidateTokenInputBody struct {
	AccessToken string `json:"access_token" doc:"The access token to authenticate"`
}

// NewEmptyValidateTokenInputBody creates a new instance of ValidateTokenInputBody with no initialized fields.
func NewEmptyValidateTokenInputBody() *ValidateTokenInputBody {
	return &ValidateTokenInputBody{}
}

// NewValidateTokenInputBody creates a new instance of ValidateTokenInputBody with the provided access token.
func NewValidateTokenInputBody(accessToken string) *ValidateTokenInputBody {
	return &ValidateTokenInputBody{AccessToken: accessToken}
}
