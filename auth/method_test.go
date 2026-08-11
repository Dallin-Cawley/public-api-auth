package auth

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MethodTestSuite struct {
	suite.Suite
}

func (testSuite *MethodTestSuite) TestMethodsFromString_Success() {
	theMethods, err := MethodsFromString([]string{"client_secret_basic", "client_secret_post"})

	testSuite.NoError(err)
	testSuite.Equal(Methods{ClientSecretBasic, ClientSecretPost}, theMethods)
}

func (testSuite *MethodTestSuite) TestMethodsFromString_InvalidMethod() {
	_, err := MethodsFromString([]string{"invalid"})
	testSuite.ErrorContains(err, "unable to parse methods")
}

func (testSuite *MethodTestSuite) TestStrings_Success() {
	testSuite.Equal(
		[]string{"client_secret_basic", "client_secret_post"},
		Methods{ClientSecretBasic, ClientSecretPost}.Strings(),
	)
}

func (testSuite *MethodTestSuite) Test_MakeMethod_ClientSecretBasic_Success() {
	theMethod, err := MakeMethod("client_secret_basic")

	testSuite.NoError(err)
	testSuite.Equal(ClientSecretBasic, theMethod)
}

func (testSuite *MethodTestSuite) Test_MakeMethod_ClientSecretPost_Success() {
	theMethod, err := MakeMethod("client_secret_post")

	testSuite.NoError(err)
	testSuite.Equal(ClientSecretPost, theMethod)
}

func (testSuite *MethodTestSuite) Test_MakeMethod_None_Success() {
	theMethod, err := MakeMethod("none")

	testSuite.NoError(err)
	testSuite.Equal(None, theMethod)
}

func (testSuite *MethodTestSuite) Test_MakeMethod_InvalidMethod() {
	theMethod, err := MakeMethod("invalid")

	testSuite.ErrorContains(err, "invalid token endpoint auth method")
	testSuite.Equal(MethodUnknown, theMethod)
}

func (testSuite *MethodTestSuite) Test_String_ClientSecretBasic_Success() {
	testSuite.Equal("client_secret_basic", ClientSecretBasic.String())
}

func (testSuite *MethodTestSuite) Test_String_ClientSecretPost_Success() {
	testSuite.Equal("client_secret_post", ClientSecretPost.String())
}

func (testSuite *MethodTestSuite) Test_MarshalJSON_Success() {
	methodBytes, err := json.Marshal(ClientSecretBasic)

	testSuite.NoError(err)
	testSuite.Equal(fmt.Sprintf(`"%s"`, ClientSecretBasic.String()), string(methodBytes))
}

func (testSuite *MethodTestSuite) Test_UnmarshalJSON_Success() {
	bytes, err := json.Marshal(ClientSecretPost)
	testSuite.NoError(err)

	var method Method
	err = json.Unmarshal(bytes, &method)

	testSuite.NoError(err)
	testSuite.Equal(ClientSecretPost, method)
}

func (testSuite *MethodTestSuite) Test_UnmarshalJSON_InvalidMethod() {
	var method Method
	err := json.Unmarshal([]byte(`"invalid method"`), &method)

	testSuite.ErrorContains(err, "invalid token endpoint auth method")
	testSuite.Equal(MethodUnknown, method)
}

func Test_RunMethodTestSuite(t *testing.T) {
	suite.Run(t, new(MethodTestSuite))
}
