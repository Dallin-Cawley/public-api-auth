package input

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RotateTestSuite struct {
	suite.Suite
}

func (testSuite *RotateTestSuite) Test_NewRotateCredentialsInputBody() {
	body := NewRotateCredentialsInputBody(RotateRegistrationAccessToken, RotateClientSecret)

	testSuite.Equal([]RotationTarget{RotateRegistrationAccessToken, RotateClientSecret}, body.Rotate)
}

func (testSuite *RotateTestSuite) Test_Valid() {
	for _, target := range RotationTargets {
		testSuite.True(target.Valid(), target)
	}

	testSuite.False(RotationTarget("").Valid())
	testSuite.False(RotationTarget("client_id").Valid())
	testSuite.False(RotationTarget("CLIENT_SECRET").Valid())
}

func (testSuite *RotateTestSuite) Test_JSON() {
	encoded, err := json.Marshal(NewRotateCredentialsInputBody(RotateClientSecret))
	testSuite.Require().NoError(err)
	testSuite.JSONEq(`{"rotate":["client_secret"]}`, string(encoded))

	var decoded RotateCredentialsInputBody
	testSuite.Require().NoError(json.Unmarshal([]byte(`{"rotate":["registration_access_token","client_secret"]}`), &decoded))
	testSuite.Equal([]RotationTarget{RotateRegistrationAccessToken, RotateClientSecret}, decoded.Rotate)
}

func Test_RunRotateTestSuite(t *testing.T) {
	suite.Run(t, new(RotateTestSuite))
}
