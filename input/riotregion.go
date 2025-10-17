package input

import "git.championtourney.com/championtourney/public/riot"

// UpdateRiotRegionInputBody represents the request body for updating Riot-supported tournament regions.
type UpdateRiotRegionInputBody struct {
	Regions []string `json:"regions" enum:"BR,EUNE,EUW,JP,LAN,LAS,NA,OCE,PBE,RU,TR,KR,PH,SG,TH,TW,VN" doc:"The regions the application intends to host tournaments in. These will be appended to the existing list of regions."`
}

// NewEmptyUpdateRiotRegionInputBody creates a new input.UpdateRiotRegionInputBody with empty values.
func NewEmptyUpdateRiotRegionInputBody() *UpdateRiotRegionInputBody {
	return &UpdateRiotRegionInputBody{}
}

// NewUpdateRiotRegionInputBody creates a new input.UpdateRiotRegionInputBody with the provided list of Riot-supported tournament regions.
func NewUpdateRiotRegionInputBody(regions []riot.Region) *UpdateRiotRegionInputBody {
	var regionStrings []string

	for _, region := range regions {
		regionStrings = append(regionStrings, region.String())
	}

	return &UpdateRiotRegionInputBody{Regions: regionStrings}
}

func (body *UpdateRiotRegionInputBody) GetRegions() []riot.Region {
	var regions []riot.Region

	for _, regionStr := range body.Regions {
		region := riot.MakeRegion(regionStr)

		regions = append(regions, region)
	}

	return regions
}
