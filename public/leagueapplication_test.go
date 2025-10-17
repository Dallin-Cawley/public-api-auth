package public

import (
	"testing"

	"git.championtourney.com/championtourney/public/integration"
	"git.championtourney.com/championtourney/public/riot"

	"github.com/stretchr/testify/suite"
)

type LeagueApplicationTestSuite struct {
	suite.Suite
	providerIDs map[riot.Region]int
}

func (testSuite *LeagueApplicationTestSuite) SetupTest() {
	testSuite.providerIDs = map[riot.Region]int{
		riot.Brazil:            100,
		riot.EuropeNorthEast:   101,
		riot.EuropeWest:        102,
		riot.Japan:             103,
		riot.LatinAmericaNorth: 104,
		riot.LatinAmericaSouth: 105,
		riot.NorthAmerica:      106,
		riot.Oceania:           107,
		riot.Pbe:               108,
		riot.Russia:            109,
		riot.Turkey:            110,
		riot.Korea:             111,
		riot.Philippines:       112,
		riot.Singapore:         113,
		riot.Thailand:          114,
		riot.Taiwan:            115,
		riot.Vietnam:           116,
	}
}

func (testSuite *LeagueApplicationTestSuite) TestNewLeagueApplication_WithAllProviderIDs() {
	applicationID := 12345

	result := NewLeagueApplication(applicationID, testSuite.providerIDs)

	testSuite.NotNil(result)
	testSuite.Equal(applicationID, result.ApplicationID)
	testSuite.Equal(100, result.BRProviderID)
	testSuite.Equal(101, result.EUNEProviderID)
	testSuite.Equal(102, result.EUWProviderID)
	testSuite.Equal(103, result.JPProviderID)
	testSuite.Equal(104, result.LANProviderID)
	testSuite.Equal(105, result.LASProviderID)
	testSuite.Equal(106, result.NAProviderID)
	testSuite.Equal(107, result.OCEProviderID)
	testSuite.Equal(108, result.PBEProviderID)
	testSuite.Equal(109, result.RUProviderID)
	testSuite.Equal(110, result.TRProviderID)
	testSuite.Equal(111, result.KRProviderID)
	testSuite.Equal(112, result.PHProviderID)
	testSuite.Equal(113, result.SGProviderID)
	testSuite.Equal(114, result.THProviderID)
	testSuite.Equal(115, result.TWProviderID)
	testSuite.Equal(116, result.VNProviderID)
	testSuite.Equal(integration.LeagueOfLegends.String(), result.IntegrationType)
}

func (testSuite *LeagueApplicationTestSuite) TestNewLeagueApplication_WithEmptyProviderIDs() {
	applicationID := 67890
	testSuite.providerIDs = map[riot.Region]int{}

	result := NewLeagueApplication(applicationID, testSuite.providerIDs)

	testSuite.NotNil(result)
	testSuite.Equal(applicationID, result.ApplicationID)
	testSuite.Equal(0, result.BRProviderID)
	testSuite.Equal(0, result.EUNEProviderID)
	testSuite.Equal(0, result.EUWProviderID)
	testSuite.Equal(0, result.JPProviderID)
	testSuite.Equal(0, result.LANProviderID)
	testSuite.Equal(0, result.LASProviderID)
	testSuite.Equal(0, result.NAProviderID)
	testSuite.Equal(0, result.OCEProviderID)
	testSuite.Equal(0, result.PBEProviderID)
	testSuite.Equal(0, result.RUProviderID)
	testSuite.Equal(0, result.TRProviderID)
	testSuite.Equal(0, result.KRProviderID)
	testSuite.Equal(0, result.PHProviderID)
	testSuite.Equal(0, result.SGProviderID)
	testSuite.Equal(0, result.THProviderID)
	testSuite.Equal(0, result.TWProviderID)
	testSuite.Equal(0, result.VNProviderID)
	testSuite.Equal(integration.LeagueOfLegends.String(), result.IntegrationType)
}

func (testSuite *LeagueApplicationTestSuite) TestNewLeagueApplication_WithPartialProviderIDs() {
	applicationID := 24680
	testSuite.providerIDs = map[riot.Region]int{
		riot.NorthAmerica: 200,
		riot.EuropeWest:   201,
		riot.Korea:        202,
	}

	result := NewLeagueApplication(applicationID, testSuite.providerIDs)

	testSuite.NotNil(result)
	testSuite.Equal(applicationID, result.ApplicationID)
	testSuite.Equal(0, result.BRProviderID)
	testSuite.Equal(0, result.EUNEProviderID)
	testSuite.Equal(201, result.EUWProviderID)
	testSuite.Equal(0, result.JPProviderID)
	testSuite.Equal(0, result.LANProviderID)
	testSuite.Equal(0, result.LASProviderID)
	testSuite.Equal(200, result.NAProviderID)
	testSuite.Equal(0, result.OCEProviderID)
	testSuite.Equal(0, result.PBEProviderID)
	testSuite.Equal(0, result.RUProviderID)
	testSuite.Equal(0, result.TRProviderID)
	testSuite.Equal(202, result.KRProviderID)
	testSuite.Equal(0, result.PHProviderID)
	testSuite.Equal(0, result.SGProviderID)
	testSuite.Equal(0, result.THProviderID)
	testSuite.Equal(0, result.TWProviderID)
	testSuite.Equal(0, result.VNProviderID)
	testSuite.Equal(integration.LeagueOfLegends.String(), result.IntegrationType)
}

func (testSuite *LeagueApplicationTestSuite) TestNewLeagueApplication_WithNilProviderIDs() {
	applicationID := 13579
	testSuite.providerIDs = nil

	result := NewLeagueApplication(applicationID, testSuite.providerIDs)

	testSuite.NotNil(result)
	testSuite.Equal(applicationID, result.ApplicationID)
	testSuite.Equal(0, result.BRProviderID)
	testSuite.Equal(0, result.EUNEProviderID)
	testSuite.Equal(0, result.EUWProviderID)
	testSuite.Equal(0, result.JPProviderID)
	testSuite.Equal(0, result.LANProviderID)
	testSuite.Equal(0, result.LASProviderID)
	testSuite.Equal(0, result.NAProviderID)
	testSuite.Equal(0, result.OCEProviderID)
	testSuite.Equal(0, result.PBEProviderID)
	testSuite.Equal(0, result.RUProviderID)
	testSuite.Equal(0, result.TRProviderID)
	testSuite.Equal(0, result.KRProviderID)
	testSuite.Equal(0, result.PHProviderID)
	testSuite.Equal(0, result.SGProviderID)
	testSuite.Equal(0, result.THProviderID)
	testSuite.Equal(0, result.TWProviderID)
	testSuite.Equal(0, result.VNProviderID)
	testSuite.Equal(integration.LeagueOfLegends.String(), result.IntegrationType)
}

func (testSuite *LeagueApplicationTestSuite) TestNewLeagueApplication_WithZeroApplicationID() {
	applicationID := 0
	testSuite.providerIDs = map[riot.Region]int{
		riot.NorthAmerica: 300,
	}

	result := NewLeagueApplication(applicationID, testSuite.providerIDs)

	testSuite.NotNil(result)
	testSuite.Equal(0, result.ApplicationID)
	testSuite.Equal(300, result.NAProviderID)
	testSuite.Equal(integration.LeagueOfLegends.String(), result.IntegrationType)
}

func (testSuite *LeagueApplicationTestSuite) TestNewLeagueApplication_WithNegativeApplicationID() {
	applicationID := -1
	testSuite.providerIDs = map[riot.Region]int{
		riot.EuropeWest: 400,
	}

	result := NewLeagueApplication(applicationID, testSuite.providerIDs)

	testSuite.NotNil(result)
	testSuite.Equal(-1, result.ApplicationID)
	testSuite.Equal(400, result.EUWProviderID)
	testSuite.Equal(integration.LeagueOfLegends.String(), result.IntegrationType)
}

func (testSuite *LeagueApplicationTestSuite) TestGetID_Success() {
	result := NewLeagueApplication(12345, testSuite.providerIDs)

	testSuite.Equal(-1, result.GetID())
}

func (testSuite *LeagueApplicationTestSuite) TestGetType_Success() {
	result := NewLeagueApplication(12345, testSuite.providerIDs)

	testSuite.Equal(integration.LeagueOfLegends, result.GetType())
}

func (testSuite *LeagueApplicationTestSuite) TestGetProviderID_AllRegions() {
	leagueApplication := NewLeagueApplication(12345, testSuite.providerIDs)

	testSuite.Equal(100, leagueApplication.GetProviderID(riot.Brazil))
	testSuite.Equal(101, leagueApplication.GetProviderID(riot.EuropeNorthEast))
	testSuite.Equal(102, leagueApplication.GetProviderID(riot.EuropeWest))
	testSuite.Equal(103, leagueApplication.GetProviderID(riot.Japan))
	testSuite.Equal(104, leagueApplication.GetProviderID(riot.LatinAmericaNorth))
	testSuite.Equal(105, leagueApplication.GetProviderID(riot.LatinAmericaSouth))
	testSuite.Equal(106, leagueApplication.GetProviderID(riot.NorthAmerica))
	testSuite.Equal(107, leagueApplication.GetProviderID(riot.Oceania))
	testSuite.Equal(108, leagueApplication.GetProviderID(riot.Pbe))
	testSuite.Equal(109, leagueApplication.GetProviderID(riot.Russia))
	testSuite.Equal(110, leagueApplication.GetProviderID(riot.Turkey))
	testSuite.Equal(111, leagueApplication.GetProviderID(riot.Korea))
	testSuite.Equal(112, leagueApplication.GetProviderID(riot.Philippines))
	testSuite.Equal(113, leagueApplication.GetProviderID(riot.Singapore))
	testSuite.Equal(114, leagueApplication.GetProviderID(riot.Thailand))
	testSuite.Equal(115, leagueApplication.GetProviderID(riot.Taiwan))
	testSuite.Equal(116, leagueApplication.GetProviderID(riot.Vietnam))
}

func (testSuite *LeagueApplicationTestSuite) TestGetProviderID_DefaultRegion() {
	testSuite.providerIDs = map[riot.Region]int{
		riot.NorthAmerica: 500,
	}
	leagueApplication := NewLeagueApplication(12345, testSuite.providerIDs)

	testSuite.Equal(500, leagueApplication.GetProviderID(riot.Region(-1)))
}

func (testSuite *LeagueApplicationTestSuite) TestGetProviderID_EmptyProviderIDs() {
	testSuite.providerIDs = map[riot.Region]int{}
	leagueApplication := NewLeagueApplication(12345, testSuite.providerIDs)

	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Brazil))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.EuropeNorthEast))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.EuropeWest))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Japan))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.LatinAmericaNorth))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.LatinAmericaSouth))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.NorthAmerica))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Oceania))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Pbe))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Russia))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Turkey))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Korea))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Philippines))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Singapore))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Thailand))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Taiwan))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Vietnam))
}

func (testSuite *LeagueApplicationTestSuite) TestGetProviderID_PartialProviderIDs() {
	testSuite.providerIDs = map[riot.Region]int{
		riot.NorthAmerica: 600,
		riot.EuropeWest:   601,
		riot.Korea:        602,
	}
	leagueApplication := NewLeagueApplication(12345, testSuite.providerIDs)

	testSuite.Equal(600, leagueApplication.GetProviderID(riot.NorthAmerica))
	testSuite.Equal(601, leagueApplication.GetProviderID(riot.EuropeWest))
	testSuite.Equal(602, leagueApplication.GetProviderID(riot.Korea))

	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Brazil))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Japan))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Russia))
}

func (testSuite *LeagueApplicationTestSuite) TestGetProviderID_NilProviderIDs() {
	testSuite.providerIDs = nil
	leagueApplication := NewLeagueApplication(12345, testSuite.providerIDs)

	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Brazil))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.EuropeNorthEast))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.EuropeWest))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Japan))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.LatinAmericaNorth))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.LatinAmericaSouth))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.NorthAmerica))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Oceania))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Pbe))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Russia))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Turkey))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Korea))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Philippines))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Singapore))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Thailand))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Taiwan))
	testSuite.Equal(0, leagueApplication.GetProviderID(riot.Vietnam))
}

func Test_LeagueApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(LeagueApplicationTestSuite))
}
