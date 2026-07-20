package grant

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
	ClientCredentials Type = iota
	AuthorizationCode
	Unsupported
)

var (
	typeStringMap = map[string]Type{
		ClientCredentials.String(): ClientCredentials,
		AuthorizationCode.String(): AuthorizationCode,
		Unsupported.String():       Unsupported,
	}
)

// MakeType maps the provided string to its grant.Type
func MakeType(grantTypeStr string) (Type, error) {
	if grantType, exists := typeStringMap[grantTypeStr]; exists {
		return grantType, nil
	}

	return -1, fmt.Errorf("invalid grant type [ %s ]", grantTypeStr)
}

func (grantType Type) String() string {
	return []string{"client_credentials", "authorization_code", "unsupported"}[grantType]
}

// MarshalJSON marshals the grant.Type on the value receiver to ensure that usages of grant.Type are
// serialized properly
func (grantType Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(grantType.String())
}

// UnmarshalJSON unmarshals the grant.Type on the pointer receiver to ensure that when passed to json#Unmarshal,
// the pointer to the grantType are properly deserialized.
func (grantType *Type) UnmarshalJSON(data []byte) (err error) {
	if *grantType, err = MakeType(strings.ReplaceAll(string(data), `"`, ``)); err != nil {
		return
	}

	return
}
