package public

import "git.championtourney.com/championtourney/public/action"

// OneTimeCodeMetadata represents metadata for a one-time code, including its action type and associated meta-data.
type OneTimeCodeMetadata struct {
	Action   string         `json:"action" doc:"The action that the one time code is for"`
	MetaData map[string]any `json:"metadata" doc:"The meta data associated with the one time code. This is the data that was sent when creating the one time code"`
}

// NewOneTimeCodeMetadata creates a new instance of OneTimeCodeMetadata with the specified action and meta-data.
func NewOneTimeCodeMetadata(action string, metaData map[string]any) *OneTimeCodeMetadata {
	return &OneTimeCodeMetadata{
		Action:   action,
		MetaData: metaData,
	}
}

// GetAction converts the public.OneTimeCodeMetaData#Action to its action.Action counterpart
func (body *OneTimeCodeMetadata) GetAction() (action.Action, error) {
	return action.MakeAction(body.Action)
}
