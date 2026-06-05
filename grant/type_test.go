package grant

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TypeTestSuite struct {
	suite.Suite
}

func (testSuite *TypeTestSuite) Test_MakeType_ClientCredentials_Success() {
	theGrant, err := MakeType("client_credentials")

	testSuite.NoError(err)
	testSuite.Equal(ClientCredentials, theGrant)
}

func (testSuite *TypeTestSuite) Test_MakeType_AuthorizationCode_Success() {
	theGrant, err := MakeType("authorization_code")

	testSuite.NoError(err)
	testSuite.Equal(AuthorizationCode, theGrant)
}

func (testSuite *TypeTestSuite) Test_MakeType_Unsupported_Success() {
	theGrant, err := MakeType("unsupported")

	testSuite.NoError(err)
	testSuite.Equal(Unsupported, theGrant)
}

func (testSuite *TypeTestSuite) Test_MakeType_InvalidGrantType() {
	theGrant, err := MakeType("invalid")

	testSuite.ErrorContains(err, "invalid grant type")
	testSuite.Equal(Type(-1), theGrant)
}

func (testSuite *TypeTestSuite) Test_String_ClientCredentials_Success() {
	testSuite.Equal("client_credentials", ClientCredentials.String())
}

func (testSuite *TypeTestSuite) Test_String_AuthorizationCode_Success() {
	testSuite.Equal("authorization_code", AuthorizationCode.String())
}

func (testSuite *TypeTestSuite) Test_String_Unsupported_Success() {
	testSuite.Equal("unsupported", Unsupported.String())
}

func (testSuite *TypeTestSuite) Test_MarshalJSON_Success() {
	typeBytes, err := json.Marshal(AuthorizationCode)

	testSuite.NoError(err)
	testSuite.Equal(fmt.Sprintf(`"%s"`, AuthorizationCode.String()), string(typeBytes))
}

func (testSuite *TypeTestSuite) Test_UnmarshalJSON_Success() {
	bytes, err := json.Marshal(ClientCredentials)
	testSuite.NoError(err)

	var grantType Type
	err = json.Unmarshal(bytes, &grantType)

	testSuite.NoError(err)
	testSuite.Equal(ClientCredentials, grantType)
}

func (testSuite *TypeTestSuite) Test_UnmarshalJSON_InvalidType() {
	var grantType Type
	err := json.Unmarshal([]byte(`"invalid type"`), &grantType)

	testSuite.ErrorContains(err, "invalid grant type")
	testSuite.Equal(Type(-1), grantType)
}

func Test_RunGrantTypeTestSuite(t *testing.T) {
	suite.Run(t, new(TypeTestSuite))
}
