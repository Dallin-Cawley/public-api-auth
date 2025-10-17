package grant

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Type int

const (
	ClientCredentials Type = iota
	AuthorizationCode
)

var (
	typeStringMap = map[string]Type{
		ClientCredentials.String(): ClientCredentials,
		AuthorizationCode.String(): AuthorizationCode,
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
	return []string{"client_credentials", "authorization_code"}[grantType]
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
