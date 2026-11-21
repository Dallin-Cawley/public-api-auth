package output

import "time"

// CreateTokenOutputBody provides the structure of the output body in the POST /token request
type CreateTokenOutputBody struct {
	AccessToken   string `json:"access_token" doc:"The JWT access token"`
	AccessExpires string `json:"access_expires" doc:"When the access token expires formatted according to the RFC3339 time format"`
	// ExpiresIn is the access token's lifetime in seconds (RFC 6749 Section 5.1), the form OAuth clients read; it
	// names the same moment as AccessExpires (CHAMP-209).
	ExpiresIn int64  `json:"expires_in" doc:"The lifetime of the access token in seconds (RFC 6749 Section 5.1)"`
	TokenType string `json:"token_type" enum:"Bearer" doc:"The type of token"`
	// RefreshToken is set when the grant issues one (CHAMP-180): an authorization_code or refresh_token grant of a
	// client registered for refresh_token. It replaces the one that was used (rotation).
	RefreshToken string `json:"refresh_token,omitempty" doc:"The refresh token, when one is issued (RFC 6749 Section 5.1)"`
	// Scope is the scope of the access token, set when it differs from the requested one (RFC 6749 Section 5.1).
	Scope string `json:"scope,omitempty" doc:"The scope of the access token, when it differs from the requested scope"`
}

// NewCreateTokenOutputBody creates the output body for the POST /token request from the access token, the moment it
// expires and its lifetime (expires minus the moment it was issued). access_expires is expires in UTC, RFC 3339;
// expires_in is the lifetime in whole seconds.
func NewCreateTokenOutputBody(accessToken string, expires time.Time, lifetime time.Duration) *CreateTokenOutputBody {
	return &CreateTokenOutputBody{
		AccessToken:   accessToken,
		AccessExpires: expires.UTC().Format(time.RFC3339),
		ExpiresIn:     int64(lifetime / time.Second),
		TokenType:     "Bearer",
	}
}

// WithRefreshToken sets the refresh token issued with the access token.
func (body *CreateTokenOutputBody) WithRefreshToken(refreshToken string) *CreateTokenOutputBody {
	body.RefreshToken = refreshToken
	return body
}

// WithScope sets the scope of the access token.
func (body *CreateTokenOutputBody) WithScope(scope string) *CreateTokenOutputBody {
	body.Scope = scope
	return body
}
