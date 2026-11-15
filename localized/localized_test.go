package localized

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LocalizedTestSuite struct {
	suite.Suite
}

func (testSuite *LocalizedTestSuite) TestSplit() {
	for member, want := range map[string][3]any{
		"client_name#ja-Jpan-JP": {"client_name", "ja-Jpan-JP", true},
		"policy_uri#fr":          {"policy_uri", "fr", true},
		"client_name":            {"", "", false},
		"client_name#":           {"", "", false},
		"jwks_uri#fr":            {"", "", false},
		"software_id#fr":         {"", "", false},
	} {
		field, tag, ok := Split(member)
		testSuite.Equal(want, [3]any{field, tag, ok}, member)
	}
}

// DCR-013: the tag is checked against the RFC 5646 Section 2.1 ABNF, in any case.
func (testSuite *LocalizedTestSuite) TestValidTag() {
	for _, tag := range []string{
		"fr", "FR", "ja-Jpan-JP", "en-US", "zh-Hant-TW", "x-private", "x-a-b", "de-CH-1901", "sl-rozaj-biske", "es-419",
		"zh-cmn-Hans-CN", "zh-min-nan", "en-a-bbb-x-a-ccc", "de-DE-u-co-phonebk", "qaa-Qaaa-QM-x-southern", "abcd",
		"abcdefgh", "i-klingon", "I-Default", "en-GB-oed", "sgn-BE-FR", "art-lojban", "hy-Latn-IT-arevela",
	} {
		testSuite.True(ValidTag(tag), tag)
	}
	for _, tag := range []string{
		"", "-fr", "fr-", "fr--CA", "fr_CA", "toolongtag", "1fr", "fr CA", "x", "a", "q-q-q", "a-DE", "fr-x", "en-a",
		"en-a-b", "de-419-DE", "ab-abc-abc-abc-abc", "en-US-u", "i-nonexistent", "fré",
	} {
		testSuite.False(ValidTag(tag), tag)
	}
}

// DCR-014: a merged member replaces one whose language tag differs only in case.
func (testSuite *LocalizedTestSuite) TestMerge() {
	testSuite.Nil(Merge(nil, nil))
	testSuite.Equal(Values{"client_name#fr": "Nom"}, Merge(nil, Values{"client_name#fr": "Nom"}))

	target := Values{"client_name#fr": "request", "client_name#de": "Name", "logo_uri#fr": "https://x/fr.png"}
	merged := Merge(target, Values{"client_name#FR": "statement", "tos_uri#en": "https://x/tos"})
	testSuite.Equal(Values{
		"client_name#FR": "statement", "client_name#de": "Name", "logo_uri#fr": "https://x/fr.png", "tos_uri#en": "https://x/tos",
	}, merged)

	// Variants of one document that differ only in case are both kept, for the caller to reject.
	testSuite.Equal(Values{"client_name#fr": "a", "client_name#FR": "b"}, Merge(nil, Values{"client_name#fr": "a", "client_name#FR": "b"}))
	testSuite.Equal(Values{"client_name#fr": "a", "client_name#FR": "b"},
		Merge(Values{"client_name#Fr": "old"}, Values{"client_name#fr": "a", "client_name#FR": "b"}))

	// The same member name simply takes the new value.
	testSuite.Equal(Values{"client_name#fr": "b"}, Merge(Values{"client_name#fr": "a"}, Values{"client_name#fr": "b"}))
}

func (testSuite *LocalizedTestSuite) TestIsURI() {
	testSuite.True(IsURI("logo_uri"))
	testSuite.False(IsURI("client_name"))
}

func (testSuite *LocalizedTestSuite) TestCollect() {
	values, err := Collect([]byte(`{"client_name":"n","client_name#fr":"Nom","logo_uri#de":"https://x/de.png","other#fr":"x"}`))
	testSuite.Require().NoError(err)
	testSuite.Equal(Values{"client_name#fr": "Nom", "logo_uri#de": "https://x/de.png"}, values)
}

func (testSuite *LocalizedTestSuite) TestCollect_None() {
	values, err := Collect([]byte(`{"client_name":"n"}`))
	testSuite.Require().NoError(err)
	testSuite.Nil(values)
}

func (testSuite *LocalizedTestSuite) TestCollect_NotAnObject() {
	values, err := Collect([]byte(`[1]`))
	testSuite.Require().NoError(err)
	testSuite.Nil(values)
}

func (testSuite *LocalizedTestSuite) TestCollect_NonStringValue() {
	_, err := Collect([]byte(`{"client_name#fr":5}`))

	var typeErr *json.UnmarshalTypeError
	testSuite.Require().ErrorAs(err, &typeErr)
	testSuite.Equal("client_name#fr", typeErr.Field)
}

func (testSuite *LocalizedTestSuite) TestAppend() {
	data, err := Append([]byte(`{"client_name":"n"}`), Values{"tos_uri#fr": "https://x/fr", "client_name#fr": "Nom"})
	testSuite.Require().NoError(err)
	testSuite.JSONEq(`{"client_name":"n","client_name#fr":"Nom","tos_uri#fr":"https://x/fr"}`, string(data))
	testSuite.Equal(`{"client_name":"n","client_name#fr":"Nom","tos_uri#fr":"https://x/fr"}`, string(data))
}

func (testSuite *LocalizedTestSuite) TestAppend_EmptyObject() {
	data, err := Append([]byte(`{}`), Values{"client_name#fr": "Nom"})
	testSuite.Require().NoError(err)
	testSuite.Equal(`{"client_name#fr":"Nom"}`, string(data))
}

func (testSuite *LocalizedTestSuite) TestAppend_SkipsMembersThatArentTaggedFields() {
	data, err := Append([]byte(`{"client_name":"n"}`), Values{"client_id": "evil", "client_name": "evil", "jwks_uri#fr": "x"})
	testSuite.Require().NoError(err)
	testSuite.Equal(`{"client_name":"n"}`, string(data))
}

func (testSuite *LocalizedTestSuite) TestAppend_NotAnObject() {
	_, err := Append([]byte(`[]`), Values{"client_name#fr": "Nom"})
	testSuite.Error(err)
}

func TestLocalizedTestSuite(t *testing.T) {
	suite.Run(t, new(LocalizedTestSuite))
}
