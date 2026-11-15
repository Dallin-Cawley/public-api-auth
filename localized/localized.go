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

// tagPattern is the langtag and privateuse productions of the RFC 5646 Section 2.1 ABNF, matched case-insensitively
// (Section 2.1.1):
//
//	langtag    = language ["-" script] ["-" region] *("-" variant) *("-" extension) ["-" privateuse]
//	language   = 2*3ALPHA ["-" extlang] / 4ALPHA / 5*8ALPHA      extlang = 3ALPHA *2("-" 3ALPHA)
//	script     = 4ALPHA                                          region  = 2ALPHA / 3DIGIT
//	variant    = 5*8alphanum / (DIGIT 3alphanum)
//	extension  = singleton 1*("-" (2*8alphanum))                 singleton = any alphanum but "x"
//	privateuse = "x" 1*("-" (1*8alphanum))
var tagPattern = regexp.MustCompile(`(?i)^(?:` +
	`(?:[a-z]{2,3}(?:-[a-z]{3}){0,3}|[a-z]{4}|[a-z]{5,8})` + // language
	`(?:-[a-z]{4})?` + // script
	`(?:-(?:[a-z]{2}|[0-9]{3}))?` + // region
	`(?:-(?:[a-z0-9]{5,8}|[0-9][a-z0-9]{3}))*` + // variants
	`(?:-[0-9a-wy-z](?:-[a-z0-9]{2,8})+)*` + // extensions
	`(?:-x(?:-[a-z0-9]{1,8})+)?` + // private use
	`|x(?:-[a-z0-9]{1,8})+` + // a tag that is private use only
	`)$`)

// irregularTags are the grandfathered tags of RFC 5646 Section 2.1 that don't match the langtag production. The
// regular grandfathered tags (such as zh-min-nan) do.
var irregularTags = []string{
	"en-GB-oed", "i-ami", "i-bnn", "i-default", "i-enochian", "i-hak", "i-klingon", "i-lux", "i-mingo", "i-navajo",
	"i-pwn", "i-tao", "i-tay", "i-tsu", "sgn-BE-FR", "sgn-BE-NL", "sgn-CH-DE",
}

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

// ValidTag reports whether tag is a well-formed language tag as RFC 5646 Section 2.1 defines it, which RFC 7591
// Section 2.2 requires of the tags in claim names. It checks the syntax only, in any character case, not the subtag
// registry: "fr", "zh-Hant-TW" and "x-private" are well-formed, "x", "a" and "q-q-q" are not.
func ValidTag(tag string) bool {
	if tagPattern.MatchString(tag) {
		return true
	}

	return slices.ContainsFunc(irregularTags, func(irregular string) bool { return strings.EqualFold(irregular, tag) })
}

// Merge adds values to target and returns it. A language tag is compared case-insensitively (RFC 7591 Section 2.2), so
// a member of values replaces a member of target whose name differs only in case: "client_name#FR" replaces
// "client_name#fr". This is what lets the claims of a software statement take precedence over the request's own
// members (RFC 7591 Section 2). A nil target is allocated when there is something to add.
func Merge(target, values Values) Values {
	for member, value := range values {
		for existing := range target {
			if existing != member && strings.EqualFold(existing, member) {
				delete(target, existing)
			}
		}

		if target == nil {
			target = Values{}
		}
		target[member] = value
	}

	return target
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
