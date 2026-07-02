package output

import "github.com/Dallin-Cawley/public-api-auth/public"

type UpdateRiotRegionOutputBody struct {
	LeagueApplication public.LeagueApplication `json:"league_application" doc:"The updated league integration for the application"`
}

func NewUpdateRiotRegionOutputBody(leagueApplication *public.LeagueApplication) *UpdateRiotRegionOutputBody {
	return &UpdateRiotRegionOutputBody{
		LeagueApplication: *leagueApplication,
	}
}
