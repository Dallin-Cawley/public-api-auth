package response

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Types []Type

func TypesFromString(stringTypes []string) (Types, error) {
	var err error

	types := make([]Type, len(stringTypes))
	for i, stringType := range stringTypes {
		types[i], err = MakeType(stringType)
		if err != nil {
			return nil, fmt.Errorf("unable to parse types: %w", err)
		}
	}

	return types, nil
}

func (types Types) Strings() []string {
	stringTypes := make([]string, len(types))
	for i, theType := range types {
		stringTypes[i] = theType.String()
	}

	return stringTypes
}

type Type int

const (
	TypeUnknown Type = iota
	Code
	Token
)

var (
	typeStringMap = map[string]Type{
		Code.String():  Code,
		Token.String(): Token,
	}
)

// MakeType maps the provided string to its response.Type
func MakeType(responseTypeStr string) (Type, error) {
	if responseType, exists := typeStringMap[responseTypeStr]; exists {
		return responseType, nil
	}

	return TypeUnknown, fmt.Errorf("invalid response type [ %s ]", responseTypeStr)
}

func (responseType Type) String() string {
	return []string{"type_unknown", "code", "token"}[responseType]
}

// MarshalJSON marshals the response.Type on the value receiver to ensure that usages of response.Type are
// serialized properly
func (responseType Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(responseType.String())
}

// UnmarshalJSON unmarshals the response.Type on the pointer receiver to ensure that when passed to json#Unmarshal,
// the pointer to the responseType are properly deserialized.
func (responseType *Type) UnmarshalJSON(data []byte) (err error) {
	if *responseType, err = MakeType(strings.ReplaceAll(string(data), `"`, ``)); err != nil {
		return
	}

	return
}
