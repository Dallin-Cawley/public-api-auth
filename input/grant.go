package input

// ClientCredentialsInputBody represents the input sent to APIAuth when creating a new model.Application
type ClientCredentialsInputBody struct {
	ApplicationName string  `json:"application_name" doc:"The name of the application"`
	CallbackURL     *string `json:"callback_url" doc:"The callback URL for the application. This will be used when emitting events related to tournaments created by the application"`
}

// NewEmptyClientCredentialsInputBody creates a new empty input.ClientCredentialsInputBody
func NewEmptyClientCredentialsInputBody() *ClientCredentialsInputBody {
	return &ClientCredentialsInputBody{}
}

// NewClientCredentialsInputBody creates an input.ClientCredentialsInputBody
func NewClientCredentialsInputBody(applicationName string, callbackURL string) *ClientCredentialsInputBody {
	return &ClientCredentialsInputBody{ApplicationName: applicationName, CallbackURL: &callbackURL}
}
