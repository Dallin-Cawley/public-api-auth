package output

// ValidateOutputBody represents the output body of the api-auth service when verifying a
// jwt token
type ValidateOutputBody struct {
	Subject   string   `json:"sub"`
	IssuedAt  string   `json:"iat"`
	ExpiresAt string   `json:"exp"`
	JWTID     string   `json:"jti"`
	Scopes    []string `json:"scope"`
}

// NewValidateOutputBody creates a new ValidateOutputBody
func NewValidateOutputBody(subject, issuedAt, expiresAt, jwtID string, scopes []string) *ValidateOutputBody {
	return &ValidateOutputBody{
		Subject:   subject,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
		JWTID:     jwtID,
		Scopes:    scopes,
	}
}
