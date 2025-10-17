package output

import "git.championtourney.com/championtourney/public-api-auth/public"

type GetApplicationOutputBody struct {
	Application public.Application `json:"application" doc:"The application associated with the provided applicationID"`
}

func NewGetApplicationOutputBody(application public.Application) *GetApplicationOutputBody {
	return &GetApplicationOutputBody{
		Application: application,
	}
}
