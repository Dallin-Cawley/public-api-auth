package input

type CreateScopesInputBody struct {
	Scopes []string `json:"scopes" schema:"scopes"`
}

func NewCreateScopesInputBody(scopes []string) *CreateScopesInputBody {
	return &CreateScopesInputBody{Scopes: scopes}
}
