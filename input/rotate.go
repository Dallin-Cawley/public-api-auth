package input

// RotationTarget names a credential value that POST /credentials/{client_id}/rotate replaces.
type RotationTarget string

const (
	// RotateRegistrationAccessToken replaces the client's registration access token (RFC 7592 Section 3).
	RotateRegistrationAccessToken RotationTarget = "registration_access_token"
	// RotateClientSecret replaces the client's client_secret.
	RotateClientSecret RotationTarget = "client_secret"
)

// RotationTargets lists every value RotateCredentialsInputBody.Rotate accepts.
var RotationTargets = []RotationTarget{RotateRegistrationAccessToken, RotateClientSecret}

// Valid reports whether target is one of RotationTargets.
func (target RotationTarget) Valid() bool {
	return target == RotateRegistrationAccessToken || target == RotateClientSecret
}

// RotateCredentialsInputBody is the body of POST /credentials/{client_id}/rotate. Rotate must name at least one
// value, so a caller always says what is being rotated.
type RotateCredentialsInputBody struct {
	Rotate []RotationTarget `json:"rotate" schema:"rotate" enum:"registration_access_token,client_secret" minItems:"1" doc:"The credential values to rotate"`
}

// NewRotateCredentialsInputBody creates a RotateCredentialsInputBody that rotates targets.
func NewRotateCredentialsInputBody(targets ...RotationTarget) *RotateCredentialsInputBody {
	return &RotateCredentialsInputBody{Rotate: targets}
}
