package public

import (
	"encoding/json"
	"errors"
	"strings"

	localErrors "git.championtourney.com/championtourney/errors/errors"
	"git.championtourney.com/championtourney/public/integration"
	"git.championtourney.com/championtourney/utils/utils"
)

const (
	NoIntegrationsMessage = "no integrations in application"
)

type Application struct {
	ApplicationID   int                       `json:"application_id"`
	ApplicationName string                    `json:"application_name"`
	ClientID        string                    `json:"client_id"`
	ClientSecret    string                    `json:"client_secret"`
	CallbackURL     string                    `json:"callback_url"`
	Integrations    []integration.Integration `json:"integrations"`
}

// NewApplication creates the public.Application from the courier.ApplicationCourier
func NewApplication(applicationID int, applicationName, clientID, clientSecret, callbackURl string, integrations []integration.Integration) *Application {
	return &Application{
		ApplicationID:   applicationID,
		ApplicationName: applicationName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		CallbackURL:     callbackURl,
		Integrations:    integrations,
	}
}

// GetID implements the public.Model interface to get the model.Application#ApplicationID
func (model *Application) GetID() int {
	return model.ApplicationID
}

// UnmarshalJSON will unmarshal the provided []byte into a public.Application accounting for any
// public.Integration implementation that exists on the public.Application
func (model *Application) UnmarshalJSON(data []byte) error {
	var dataMap map[string]any
	if err := json.Unmarshal(data, &dataMap); err != nil {
		return err
	}

	applicationID, err := utils.GetKeyFromMap[float64]("application_id", dataMap)
	if err != nil {
		return localErrors.Wrapf(err, "unable to get application_id from data; received [ %s ]", dataMap["application_id"])
	}

	applicationName, err := utils.GetKeyFromMap[string]("application_name", dataMap)
	if err != nil {
		return localErrors.Wrapf(err, "unable to get application_name from data; received [ %s ]", dataMap["application_name"])
	}

	clientID, err := utils.GetKeyFromMap[string]("client_id", dataMap)
	if err != nil {
		return localErrors.Wrapf(err, "unable to get client_id from data; received [ %s ]", dataMap["client_id"])
	}

	clientSecret, err := utils.GetKeyFromMap[string]("client_secret", dataMap)
	if err != nil {
		return localErrors.Wrapf(err, "unable to get client_secret from data; received [ %s ]", dataMap["client_secret"])
	}

	callbackURL, err := utils.GetKeyFromMap[string]("callback_url", dataMap)
	if err != nil {
		return localErrors.Wrapf(err, "unable to get callback_url from data; received [ %s ]", dataMap["callback_url"])
	}

	integrationInterfaces, err := utils.GetKeyFromMap[[]any]("integrations", dataMap)
	if err != nil {
		if !errors.Is(err, localErrors.NewEmptyNilError()) {
			return localErrors.Wrapf(err, "unable to get integrations from data; received [ %s ]", dataMap["integrations"])
		}

		integrationInterfaces = nil
	}

	model.ApplicationID = int(applicationID)
	model.ApplicationName = applicationName
	model.ClientID = clientID
	model.ClientSecret = clientSecret
	model.CallbackURL = callbackURL
	model.Integrations = model.getIntegrations(integrationInterfaces)

	return nil
}

func (model *Application) getIntegrations(integrationInterfaces []any) []integration.Integration {
	integrations := make([]integration.Integration, len(integrationInterfaces))

	for i, integrationInterface := range integrationInterfaces {
		integrationMap := integrationInterface.(map[string]any)

		integrations[i] = model.getIntegration(integrationMap)
	}

	return integrations
}

func (model *Application) getIntegration(integrationMap map[string]any) integration.Integration {
	integrationTypeName, _ := utils.GetKeyFromMap[string]("integration_type", integrationMap)
	integrationType, _ := integration.MakeType(integrationTypeName)

	bytes, _ := json.Marshal(integrationMap)

	var integrationModel integration.Integration
	if integrationType == integration.LeagueOfLegends {
		leagueApplication := &LeagueApplication{}
		_ = json.Unmarshal(bytes, leagueApplication)
		integrationModel = leagueApplication
	}

	return integrationModel
}

// GetApplicationIntegration retrieves the first integration with a successful cast to type T from the public.Application
func GetApplicationIntegration[T integration.Integration](application *Application) (T, error) {
	var requestedIntegration T
	var ok bool

	if len(application.Integrations) == 0 {
		return requestedIntegration, errors.New(NoIntegrationsMessage)
	}

	for _, integrationModel := range application.Integrations {
		if requestedIntegration, ok = integrationModel.(T); ok {
			break
		}
	}

	return requestedIntegration, nil
}

// IsNoApplicationIntegrationError determines if the error is an error indicating that a public.Application does
// not have any public.Integration(s).
func IsNoApplicationIntegrationError(err error) bool {
	return strings.Contains(err.Error(), NoIntegrationsMessage)
}
