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

func (testSuite *LocalizedTestSuite) TestValidTag() {
	for _, tag := range []string{"fr", "ja-Jpan-JP", "en-US", "zh-Hant-TW", "x-private"} {
		testSuite.True(ValidTag(tag), tag)
	}
	for _, tag := range []string{"", "-fr", "fr-", "fr--CA", "fr_CA", "toolongtag", "1fr", "fr CA"} {
		testSuite.False(ValidTag(tag), tag)
	}
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
