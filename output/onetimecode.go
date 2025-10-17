package output

import "git.championtourney.com/championtourney/public-api-auth/public"

type GetOneTimeCodeMetaDataOutputBody struct {
	MetaData public.OneTimeCodeMetadata `json:"meta_data" doc:"The meta data associated with the one time code"`
}

func NewGetOneTimeCodeMetaDataOutputBody(metaData *public.OneTimeCodeMetadata) *GetOneTimeCodeMetaDataOutputBody {
	return &GetOneTimeCodeMetaDataOutputBody{
		MetaData: *metaData,
	}
}

type CreateOneTimeCodeOutputBody struct {
	OneTimeCode string `json:"one_time_code" doc:"The one time code that was created"`
}

func NewCreateOneTimeCodeOutputBody(oneTimeCode string) *CreateOneTimeCodeOutputBody {
	return &CreateOneTimeCodeOutputBody{OneTimeCode: oneTimeCode}
}
