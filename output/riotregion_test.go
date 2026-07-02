package output

import (
	"testing"

	"git.championtourney.com/championtourney/public/riot"
	"github.com/Dallin-Cawley/public-api-auth/public"

	"github.com/stretchr/testify/suite"
)

type RiotRegionTestSuite struct {
	suite.Suite
}

func (testSuite *RiotRegionTestSuite) TestNewUpdateRiotRegionOutputBody_Success() {
	expectedLeagueApplication := public.NewLeagueApplication(1, make(map[riot.Region]int))

	body := NewUpdateRiotRegionOutputBody(expectedLeagueApplication)

	testSuite.Equal(*expectedLeagueApplication, body.LeagueApplication)
}

func Test_RunRiotRegionTestSuite(t *testing.T) {
	suite.Run(t, new(RiotRegionTestSuite))
}
