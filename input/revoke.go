package input

// RevokeTokenInputBody is the body of a token revocation request (RFC 7009 Section 2.1), POST /revoke, sent as
// application/x-www-form-urlencoded. The client authenticates as at the token endpoint.
type RevokeTokenInputBody struct {
	Token         string `json:"token" schema:"token" doc:"The token to revoke"`
	TokenTypeHint string `json:"token_type_hint" schema:"token_type_hint" enum:"refresh_token,access_token" doc:"A hint about the type of the token (RFC 7009 Section 2.1)"`
	ClientID      string `json:"client_id" schema:"client_id" doc:"The client ID, when the client doesn't authenticate with HTTP Basic"`
	ClientSecret  string `json:"client_secret" schema:"client_secret" doc:"The client secret, for client_secret_post"`
}
