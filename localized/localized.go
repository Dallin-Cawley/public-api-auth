// Package localized handles the language-tagged variants of human-readable client metadata (RFC 7591 Section 2.2),
// such as client_name#ja-Jpan-JP. A variant is a top-level JSON member named "<field>#<BCP 47 language tag>".
package localized

import (
	"bytes"
	"encoding/json"
	"reflect"
	"regexp"
	"slices"
	"strings"
)

// Fields are the human-readable client metadata members that may carry a language tag (RFC 7591 Section 2.2).
var Fields = []string{"client_name", "client_uri", "logo_uri", "tos_uri", "policy_uri"}

// Values maps a language-tagged member name, such as "client_name#fr", to its value.
type Values map[string]string

// tagPattern is the shape of a BCP 47 language tag (RFC 5646 Section 2.1): subtags of 1 to 8 letters and digits joined
// by hyphens, starting with a letter subtag (or the private-use "x" prefix).
var tagPattern = regexp.MustCompile(`^[A-Za-z]{1,8}(-[A-Za-z0-9]{1,8})*$`)

// Split splits member into its field and language tag. ok is false when member has no "#", the tag is empty, or the
// field isn't one of Fields, in which case member is ordinary (unknown) metadata.
func Split(member string) (field, tag string, ok bool) {
	field, tag, found := strings.Cut(member, "#")
	if !found || tag == "" || !slices.Contains(Fields, field) {
		return "", "", false
	}

	return field, tag, true
}

// IsURI reports whether field holds a URI, so its variants must be validated like the untagged value.
func IsURI(field string) bool {
	return strings.HasSuffix(field, "_uri")
}

// ValidTag reports whether tag is a well-formed BCP 47 language tag. It checks the syntax only, not the registry.
func ValidTag(tag string) bool {
	return tagPattern.MatchString(tag)
}

// Collect returns the language-tagged members of the JSON object in data, or nil when there are none. A tagged
// member whose value isn't a string is reported as a *json.UnmarshalTypeError naming the member. data that isn't a JSON
// object yields no values; decoding it into the body reports that error.
func Collect(data []byte) (Values, error) {
	var members map[string]json.RawMessage
	if err := json.Unmarshal(data, &members); err != nil {
		return nil, nil //nolint:nilerr // the struct decode reports a body that isn't an object
	}

	var values Values
	for member, raw := range members {
		if _, _, ok := Split(member); !ok {
			continue
		}

		var value string
		if err := json.Unmarshal(raw, &value); err != nil {
			return nil, &json.UnmarshalTypeError{Value: "non-string", Type: reflect.TypeFor[string](), Field: member}
		}

		if values == nil {
			values = Values{}
		}
		values[member] = value
	}

	return values, nil
}

// Append adds values to the JSON object in data as top-level members, sorted by name. Members whose name isn't a
// language-tagged field (see Split) are skipped, so values can't overwrite or inject other metadata.
func Append(data []byte, values Values) ([]byte, error) {
	members := make([]string, 0, len(values))
	for member := range values {
		if _, _, ok := Split(member); ok {
			members = append(members, member)
		}
	}
	if len(members) == 0 {
		return data, nil
	}
	slices.Sort(members)

	trimmed := bytes.TrimRight(data, " \t\r\n")
	if len(trimmed) < 2 || trimmed[len(trimmed)-1] != '}' {
		return nil, &json.UnsupportedValueError{Str: "localized values can only be added to a JSON object"}
	}

	var buffer bytes.Buffer
	buffer.Write(trimmed[:len(trimmed)-1])
	empty := len(bytes.TrimSpace(trimmed[1:len(trimmed)-1])) == 0

	for i, member := range members {
		if i > 0 || !empty {
			buffer.WriteByte(',')
		}

		name, err := json.Marshal(member)
		if err != nil {
			return nil, err
		}
		value, err := json.Marshal(values[member])
		if err != nil {
			return nil, err
		}

		buffer.Write(name)
		buffer.WriteByte(':')
		buffer.Write(value)
	}
	buffer.WriteByte('}')

	return buffer.Bytes(), nil
}
