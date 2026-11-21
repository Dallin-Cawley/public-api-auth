package output

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TokenTestSuite struct {
	suite.Suite
}

func (testSuite *TokenTestSuite) TestNewCreateTokenOutputBody_Success() {
	expires := time.Date(2026, 6, 4, 8, 0, 0, 0, time.FixedZone("MST", -7*60*60))

	outputBody := NewCreateTokenOutputBody("some access token", expires, 15*time.Minute)

	testSuite.Equal("some access token", outputBody.AccessToken)
	testSuite.Equal("2026-06-04T15:00:00Z", outputBody.AccessExpires)
	testSuite.Equal(int64(900), outputBody.ExpiresIn)
	testSuite.Equal("Bearer", outputBody.TokenType)
}

func (testSuite *TokenTestSuite) TestNewCreateTokenOutputBody_TruncatesToSeconds() {
	outputBody := NewCreateTokenOutputBody("a", time.Now(), 90*time.Second+999*time.Millisecond)

	testSuite.Equal(int64(90), outputBody.ExpiresIn)
}

func (testSuite *TokenTestSuite) TestCreateTokenOutputBody_JSON() {
	expires := time.Date(2026, 6, 4, 15, 0, 0, 0, time.UTC)

	encoded, err := json.Marshal(NewCreateTokenOutputBody("a", expires, time.Hour))
	testSuite.Require().NoError(err)

	testSuite.JSONEq(`{"access_token":"a","access_expires":"2026-06-04T15:00:00Z","expires_in":3600,"token_type":"Bearer"}`,
		string(encoded))
}

func (testSuite *TokenTestSuite) TestWithRefreshTokenAndScope() {
	outputBody := NewCreateTokenOutputBody("a", time.Now(), time.Minute).WithRefreshToken("r").WithScope("read")

	testSuite.Equal("r", outputBody.RefreshToken)
	testSuite.Equal("read", outputBody.Scope)
}

func Test_RunTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
