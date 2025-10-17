package output

// CreateTokenOutputBody provides the structure of the output body in the POST /token request
type CreateTokenOutputBody struct {
	AccessToken   string `json:"access_token" doc:"The JWT access token"`
	AccessExpires string `json:"access_expires" doc:"When the access token expires formatted according to the RFC3339 time format"`
	TokenType     string `json:"token_type" enum:"Bearer" doc:"The type of token"`
}

// NewCreateTokenOutputBody creates the output body for the POST /token request
func NewCreateTokenOutputBody(accessToken, accessExpires string) *CreateTokenOutputBody {
	return &CreateTokenOutputBody{
		AccessToken:   accessToken,
		AccessExpires: accessExpires,
		TokenType:     "Bearer",
	}
}
