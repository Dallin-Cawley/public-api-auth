package public

import (
	"git.championtourney.com/championtourney/public/integration"
	"git.championtourney.com/championtourney/public/riot"
)

type LeagueApplication struct {
	ApplicationID   int    `json:"application_id" doc:"The ID of the application"`
	BRProviderID    int    `json:"br_provider_id,omitempty" doc:"The provider ID for the Brazil region"`
	EUNEProviderID  int    `json:"eune_provider_id,omitempty" doc:"The provider ID for the Europe North East region"`
	EUWProviderID   int    `json:"eu_w_provider_id,omitempty" doc:"The provider ID for the Europe West region"`
	JPProviderID    int    `json:"jp_provider_id,omitempty" doc:"The provider ID for the Japan region"`
	LANProviderID   int    `json:"lan_provider_id,omitempty" doc:"The provider ID for the Latin America North region"`
	LASProviderID   int    `json:"las_provider_id,omitempty" doc:"The provider ID for the Latin America South region"`
	NAProviderID    int    `json:"na_provider_id,omitempty" doc:"The provider ID for the North America region"`
	OCEProviderID   int    `json:"oce_provider_id,omitempty" doc:"The provider ID for the Oceana region"`
	PBEProviderID   int    `json:"pbe_provider_id,omitempty" doc:"The provider ID for the PBE server"`
	RUProviderID    int    `json:"ru_provider_id,omitempty" doc:"The provider ID for the Russia region"`
	TRProviderID    int    `json:"tr_provider_id,omitempty" doc:"The provider ID for the Turkey region"`
	KRProviderID    int    `json:"kr_provider_id,omitempty" doc:"The provider ID for the Korea region"`
	PHProviderID    int    `json:"ph_provider_id,omitempty" doc:"The provider ID for the Philippines region"`
	SGProviderID    int    `json:"sg_provider_id,omitempty" doc:"The provider ID for the Singapore region"`
	THProviderID    int    `json:"th_provider_id,omitempty" doc:"The provider ID for the Thailand region"`
	TWProviderID    int    `json:"tw_provider_id,omitempty" doc:"The provider ID for the Taiwan region"`
	VNProviderID    int    `json:"vn_provider_id,omitempty" doc:"The provider ID for the Vietnam region"`
	IntegrationType string `json:"integration_type" enum:"LEAGUE_OF_LEGENDS" doc:"The integration this information is intended for"`
}

// NewLeagueApplication creates a new public.LeagueApplication from the provided courier.LeagueApplicationCourier.
func NewLeagueApplication(applicationID int, providerIds map[riot.Region]int) *LeagueApplication {
	return &LeagueApplication{
		ApplicationID:   applicationID,
		BRProviderID:    providerIds[riot.Brazil],
		EUNEProviderID:  providerIds[riot.EuropeNorthEast],
		EUWProviderID:   providerIds[riot.EuropeWest],
		JPProviderID:    providerIds[riot.Japan],
		LANProviderID:   providerIds[riot.LatinAmericaNorth],
		LASProviderID:   providerIds[riot.LatinAmericaSouth],
		NAProviderID:    providerIds[riot.NorthAmerica],
		OCEProviderID:   providerIds[riot.Oceania],
		PBEProviderID:   providerIds[riot.Pbe],
		RUProviderID:    providerIds[riot.Russia],
		TRProviderID:    providerIds[riot.Turkey],
		KRProviderID:    providerIds[riot.Korea],
		PHProviderID:    providerIds[riot.Philippines],
		SGProviderID:    providerIds[riot.Singapore],
		THProviderID:    providerIds[riot.Thailand],
		TWProviderID:    providerIds[riot.Taiwan],
		VNProviderID:    providerIds[riot.Vietnam],
		IntegrationType: integration.LeagueOfLegends.String(),
	}

}

// GetID implements the public.Model interface. The model.LeagueApplication#LeagueApplicationTournamentID won't be
// exposed publicly.
func (model *LeagueApplication) GetID() int {
	return -1
}

// GetType returns the integration.Type of the public.Model
func (model *LeagueApplication) GetType() integration.Type {
	return integration.LeagueOfLegends
}

// GetProviderID retrieves the appropriate providerID for the given riot.Region.
func (model *LeagueApplication) GetProviderID(region riot.Region) int {
	switch region {
	case riot.Brazil:
		return model.BRProviderID
	case riot.EuropeNorthEast:
		return model.EUNEProviderID
	case riot.EuropeWest:
		return model.EUWProviderID
	case riot.Japan:
		return model.JPProviderID
	case riot.LatinAmericaNorth:
		return model.LANProviderID
	case riot.LatinAmericaSouth:
		return model.LASProviderID
	case riot.Oceania:
		return model.OCEProviderID
	case riot.Pbe:
		return model.PBEProviderID
	case riot.Russia:
		return model.RUProviderID
	case riot.Turkey:
		return model.TRProviderID
	case riot.Korea:
		return model.KRProviderID
	case riot.Philippines:
		return model.PHProviderID
	case riot.Singapore:
		return model.SGProviderID
	case riot.Thailand:
		return model.THProviderID
	case riot.Taiwan:
		return model.TWProviderID
	case riot.Vietnam:
		return model.VNProviderID
	default:
		return model.NAProviderID
	}
}
