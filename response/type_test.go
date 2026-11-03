package response

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TypeTestSuite struct {
	suite.Suite
}

func (testSuite *TypeTestSuite) TestTypesFromString_Success() {
	theResponseTypes, err := TypesFromString([]string{"code", "token"})

	testSuite.NoError(err)
	testSuite.Equal(Types{Code, Token}, theResponseTypes)
}

func (testSuite *TypeTestSuite) TestTypesFromString_InvalidResponseType() {
	_, err := TypesFromString([]string{"invalid"})
	testSuite.ErrorContains(err, "unable to parse types")
}

func (testSuite *TypeTestSuite) TestStrings_Success() {
	testSuite.Equal(
		[]string{"code", "token"},
		Types{Code, Token}.Strings(),
	)
}

func (testSuite *TypeTestSuite) Test_MakeType_Code_Success() {
	theResponse, err := MakeType("code")

	testSuite.NoError(err)
	testSuite.Equal(Code, theResponse)
}

func (testSuite *TypeTestSuite) Test_MakeType_Token_Success() {
	theResponse, err := MakeType("token")

	testSuite.NoError(err)
	testSuite.Equal(Token, theResponse)
}

func (testSuite *TypeTestSuite) Test_MakeType_InvalidResponseType() {
	theResponse, err := MakeType("invalid")

	testSuite.ErrorContains(err, "invalid response type")
	testSuite.Equal(TypeUnknown, theResponse)
}

func (testSuite *TypeTestSuite) Test_String_Code_Success() {
	testSuite.Equal("code", Code.String())
}

func (testSuite *TypeTestSuite) Test_String_Token_Success() {
	testSuite.Equal("token", Token.String())
}

func (testSuite *TypeTestSuite) Test_MarshalJSON_Success() {
	typeBytes, err := json.Marshal(Code)

	testSuite.NoError(err)
	testSuite.Equal(fmt.Sprintf(`"%s"`, Code.String()), string(typeBytes))
}

func (testSuite *TypeTestSuite) Test_UnmarshalJSON_Success() {
	bytes, err := json.Marshal(Token)
	testSuite.NoError(err)

	var responseType Type
	err = json.Unmarshal(bytes, &responseType)

	testSuite.NoError(err)
	testSuite.Equal(Token, responseType)
}

func (testSuite *TypeTestSuite) Test_UnmarshalJSON_InvalidType() {
	var responseType Type
	err := json.Unmarshal([]byte(`"invalid type"`), &responseType)

	testSuite.ErrorContains(err, "invalid response type")
	testSuite.Equal(TypeUnknown, responseType)
}

func Test_RunResponseTypeTestSuite(t *testing.T) {
	suite.Run(t, new(TypeTestSuite))
}
