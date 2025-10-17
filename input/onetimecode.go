package input

import "git.championtourney.com/championtourney/public/action"

type CreateOneTimeCodeInputBody struct {
	Action   string            `json:"action" enum:"CREATE_TOURNAMENT,EDIT_TOURNAMENT,CREATE_TEAM,EDIT_TEAM,CREATE_MEMBER,EDIT_MEMBER" doc:"The action that the one time code will authenticate for"`
	MetaData map[string]string `json:"meta_data" doc:"The meta data associated with the one time code. This will be used by the tournament engine to execute the action for the proper entity"`
}

// NewEmptyCreateOneTimeCodeInputBody creates an input.CreateOneTimeCodeInputBody with empty values
func NewEmptyCreateOneTimeCodeInputBody() *CreateOneTimeCodeInputBody {
	return &CreateOneTimeCodeInputBody{}
}

// NewCreateOneTimeCodeInputBody creates a new input.CreateOneTimeCodeInputBody with the provided meta-data
func NewCreateOneTimeCodeInputBody(action action.Action, metaData map[string]string) *CreateOneTimeCodeInputBody {
	return &CreateOneTimeCodeInputBody{
		Action:   action.String(),
		MetaData: metaData,
	}
}

// GetAction converts the input.CreateOneTimeCodeInputBody#Action to its action.Action counterpart.
func (body *CreateOneTimeCodeInputBody) GetAction() (action.Action, error) {
	return action.MakeAction(body.Action)
}
