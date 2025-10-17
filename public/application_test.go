package public

import (
	"encoding/json"
	"testing"
	"time"

	"git.championtourney.com/championtourney/public/integration"

	"github.com/stretchr/testify/suite"
)

type ApplicationTestSuite struct {
	suite.Suite
	applicationID   int
	applicationName string
	clientID        string
	clientSecret    string
	callbackURL     string
}

func (testSuite *ApplicationTestSuite) SetupTest() {
	testSuite.applicationID = 1
	testSuite.applicationName = "some application"
	testSuite.clientID = "some client id"
	testSuite.clientSecret = "some client secret"
	testSuite.callbackURL = "some callback url"
}

func (testSuite *ApplicationTestSuite) TestNewApplication_Success() {
	publicApplication := NewApplication(testSuite.applicationID, testSuite.applicationName, testSuite.clientID, testSuite.clientSecret, testSuite.callbackURL, nil)

	testSuite.Equal(testSuite.applicationID, publicApplication.ApplicationID)
	testSuite.Equal(testSuite.applicationName, publicApplication.ApplicationName)
	testSuite.Equal(testSuite.clientID, publicApplication.ClientID)
	testSuite.Equal(testSuite.clientSecret, publicApplication.ClientSecret)
	testSuite.Equal(testSuite.callbackURL, publicApplication.CallbackURL)
}

func (testSuite *ApplicationTestSuite) TestGetID_Success() {
	publicApplication := NewApplication(testSuite.applicationID, testSuite.applicationName, testSuite.clientID, testSuite.clientSecret, testSuite.callbackURL, nil)

	testSuite.Equal(testSuite.applicationID, publicApplication.GetID())
}

func (testSuite *ApplicationTestSuite) TestGetIntegration_Success() {
	publicApplication := &Application{Integrations: []integration.Integration{&LeagueApplication{}}}

	integrationModel, err := GetApplicationIntegration[*LeagueApplication](publicApplication)

	testSuite.NoError(err)
	testSuite.IsType(&LeagueApplication{}, integrationModel)
}

func (testSuite *ApplicationTestSuite) TestGetIntegration_NoIntegrations() {
	publicApplication := &Application{Integrations: []integration.Integration{}}

	integrationModel, err := GetApplicationIntegration[*LeagueApplication](publicApplication)

	testSuite.True(IsNoApplicationIntegrationError(err))
	testSuite.Nil(integrationModel)
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_Success() {
	publicApplication := &Application{
		ApplicationID:   1,
		ApplicationName: "some name",
		ClientID:        "some id",
		ClientSecret:    "some secret",
		CallbackURL:     "some callback url",
		Integrations: []integration.Integration{
			&LeagueApplication{
				IntegrationType: integration.LeagueOfLegends.String(),
			},
		},
	}

	bytes, err := json.Marshal(publicApplication)
	testSuite.NoError(err)

	unmarshalledApplication := Application{}
	err = json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.NoError(err)
	testSuite.Equal(publicApplication.ApplicationID, unmarshalledApplication.ApplicationID)
	testSuite.Equal(publicApplication.ApplicationName, unmarshalledApplication.ApplicationName)
	testSuite.Equal(publicApplication.ClientID, unmarshalledApplication.ClientID)
	testSuite.Equal(publicApplication.ClientSecret, unmarshalledApplication.ClientSecret)
	testSuite.Equal(publicApplication.CallbackURL, unmarshalledApplication.CallbackURL)
	testSuite.NotEmpty(unmarshalledApplication.Integrations)
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_Success_NoIntegration() {
	publicApplication := &Application{
		ApplicationID:   1,
		ApplicationName: "some name",
		ClientID:        "some id",
		ClientSecret:    "some secret",
		Integrations:    nil,
	}

	bytes, err := json.Marshal(publicApplication)
	testSuite.NoError(err)

	unmarshalledApplication := Application{}
	err = json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.NoError(err)
	testSuite.Equal(publicApplication.ApplicationID, unmarshalledApplication.ApplicationID)
	testSuite.Equal(publicApplication.ApplicationName, unmarshalledApplication.ApplicationName)
	testSuite.Equal(publicApplication.ClientID, unmarshalledApplication.ClientID)
	testSuite.Equal(publicApplication.ClientSecret, unmarshalledApplication.ClientSecret)
	testSuite.Empty(unmarshalledApplication.Integrations)
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_InvalidApplicationID() {
	applicationDataMap := getApplicationMap("some application")
	applicationDataMap["application_id"] = "invalid_type"

	bytes, err := json.Marshal(applicationDataMap)
	testSuite.NoError(err)

	unmarshalledApplication := Application{}
	err = json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.ErrorContains(err, "unable to get application_id from data")
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_InvalidApplicationName() {
	applicationDataMap := getApplicationMap("some application")
	applicationDataMap["application_name"] = 0

	bytes, err := json.Marshal(applicationDataMap)
	testSuite.NoError(err)

	unmarshalledApplication := Application{}
	err = json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.ErrorContains(err, "unable to get application_name from data")
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_InvalidClientID() {
	applicationDataMap := getApplicationMap("some application")
	applicationDataMap["client_id"] = 0

	bytes, err := json.Marshal(applicationDataMap)
	testSuite.NoError(err)

	unmarshalledApplication := Application{}
	err = json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.ErrorContains(err, "unable to get client_id from data")
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_InvalidClientSecret() {
	applicationDataMap := getApplicationMap("some application")
	applicationDataMap["client_secret"] = 0

	bytes, err := json.Marshal(applicationDataMap)
	testSuite.NoError(err)

	unmarshalledApplication := Application{}
	err = json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.ErrorContains(err, "unable to get client_secret from data")
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_InvalidCallbackURL() {
	applicationDataMap := getApplicationMap("some application")
	applicationDataMap["callback_url"] = 0

	bytes, err := json.Marshal(applicationDataMap)
	testSuite.NoError(err)

	unmarshalledApplication := Application{}
	err = json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.ErrorContains(err, "unable to get callback_url from data")
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_InvalidIntegrations() {
	applicationDataMap := getApplicationMap("some application")
	applicationDataMap["integrations"] = 0

	bytes, err := json.Marshal(applicationDataMap)
	testSuite.NoError(err)

	unmarshalledApplication := Application{}
	err = json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.ErrorContains(err, "unable to get integrations from data")
}

func (testSuite *ApplicationTestSuite) TestUnmarshalJSON_InvalidJSON() {
	bytes := []byte(`"some string"`)

	unmarshalledApplication := Application{}
	err := json.Unmarshal(bytes, &unmarshalledApplication)

	testSuite.ErrorContains(err, "cannot unmarshal string into Go value")
}

// Test_ApplicationTestSuite runs the ApplicationTestSuite through testify
func Test_ApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuite))
}

func getApplicationMap(applicationName string) map[string]any {
	modelMap := make(map[string]any)

	modelMap["application_id"] = int32(1)
	modelMap["application_name"] = applicationName
	modelMap["token_id"] = int32(100)
	modelMap["client_id"] = "client_id"
	modelMap["client_secret"] = "client_secret"
	modelMap["callback_url"] = "callback"
	modelMap["created_at"] = time.Now()
	modelMap["updated_at"] = time.Now()

	return modelMap
}
