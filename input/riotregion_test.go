package input

import (
	"testing"

	"git.championtourney.com/championtourney/public/riot"

	"github.com/stretchr/testify/suite"
)

type RiotRegionTestSuite struct {
	suite.Suite
}

func (testSuite *RiotRegionTestSuite) TestNewEmptyUpdateRiotRegionBody_Success() {
	body := NewEmptyUpdateRiotRegionInputBody()

	testSuite.Empty(body.Regions)
}

func (testSuite *RiotRegionTestSuite) TestNewUpdateRiotRegionBody_Success() {
	expectedRegions := []riot.Region{riot.EuropeWest}

	body := NewUpdateRiotRegionInputBody(expectedRegions)

	for i, region := range expectedRegions {
		testSuite.Equal(region.String(), body.Regions[i])
	}
}

func (testSuite *RiotRegionTestSuite) TestGetRegions_Success_Empty() {
	expectedRegions := []riot.Region{riot.EuropeWest}

	body := NewUpdateRiotRegionInputBody(expectedRegions)

	testSuite.Equal(expectedRegions, body.GetRegions())
}

func Test_RunRiotRegionTestSuite(t *testing.T) {
	suite.Run(t, new(RiotRegionTestSuite))
}
