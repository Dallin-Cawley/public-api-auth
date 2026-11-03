package auth

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Methods []Method

func MethodsFromString(stringMethods []string) (Methods, error) {
	var err error

	methods := make([]Method, len(stringMethods))
	for i, stringMethod := range stringMethods {
		methods[i], err = MakeMethod(stringMethod)
		if err != nil {
			return nil, fmt.Errorf("unable to parse methods: %w", err)
		}
	}

	return methods, nil
}

func (methods Methods) Strings() []string {
	stringMethods := make([]string, len(methods))
	for i, theMethod := range methods {
		stringMethods[i] = theMethod.String()
	}

	return stringMethods
}

type Method int

const (
	MethodUnknown Method = iota
	ClientSecretBasic
	ClientSecretPost
	None
)

var (
	methodStringMap = map[string]Method{
		MethodUnknown.String():     MethodUnknown,
		ClientSecretBasic.String(): ClientSecretBasic,
		ClientSecretPost.String():  ClientSecretPost,
		None.String():              None,
	}
)

// MakeMethod maps the provided string to its auth.Method
func MakeMethod(methodStr string) (Method, error) {
	if method, exists := methodStringMap[methodStr]; exists {
		return method, nil
	}

	return MethodUnknown, fmt.Errorf("invalid token endpoint auth method [ %s ]", methodStr)
}

func (method Method) String() string {
	return []string{
		"method_unknown",
		"client_secret_basic",
		"client_secret_post",
		"none",
	}[method]
}

// MarshalJSON marshals the auth.Method on the value receiver to ensure that usages of auth.Method are
// serialized properly
func (method Method) MarshalJSON() ([]byte, error) {
	return json.Marshal(method.String())
}

// UnmarshalJSON unmarshals the auth.Method on the pointer receiver to ensure that when passed to json#Unmarshal,
// the pointer to the method are properly deserialized.
func (method *Method) UnmarshalJSON(data []byte) (err error) {
	if *method, err = MakeMethod(strings.ReplaceAll(string(data), `"`, ``)); err != nil {
		return
	}

	return
}
