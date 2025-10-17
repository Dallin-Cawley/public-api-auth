package output

import "git.championtourney.com/championtourney/public-api-auth/public"

type UpdateRiotRegionOutputBody struct {
	LeagueApplication public.LeagueApplication `json:"league_application" doc:"The updated league integration for the application"`
}

func NewUpdateRiotRegionOutputBody(leagueApplication *public.LeagueApplication) *UpdateRiotRegionOutputBody {
	return &UpdateRiotRegionOutputBody{
		LeagueApplication: *leagueApplication,
	}
}
