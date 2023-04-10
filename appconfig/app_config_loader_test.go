package appconfig

import (
	"encoding/json"
	"fileops"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadAppConfigSuccess(test *testing.T) {
	fileOpsMock := fileops.FileOpsMock{}
	appConfigLoader := AppConfigLoader{&fileOpsMock}

	appConfig := AppConfig{}
	appConfig.Execution.ThreadPoolSize = 1000
	configJson, _ := json.Marshal(appConfig)

	fileOpsMock.On("ReadFile", "config.json").Return(string(configJson)).Once()

	appConfigResult, err := appConfigLoader.LoadAppConfig("config.json")

	fileOpsMock.AssertExpectations(test)

	assert.NoError(test, err, "TestLoadAppConfigSuccess Failed: Unable to load config file")
	assert.Equal(test, 1000, appConfigResult.Execution.ThreadPoolSize, "LoadAppConfig Failed: Wrong ThreadPoolSize")
}

func TestLoadAppConfigFailed(test *testing.T) {
	fileOpsMock := fileops.FileOpsMock{}
	appConfigLoader := AppConfigLoader{&fileOpsMock}

	fileOpsMock.On("ReadFile", "config.json").Return("").Once()

	_, err := appConfigLoader.LoadAppConfig("config.json")

	fileOpsMock.AssertExpectations(test)

	assert.Error(test, err, "TestLoadAppConfigFailed Failed")
}

func TestLoadAppConfigFailedDeserialization(test *testing.T) {
	fileOpsMock := fileops.FileOpsMock{}
	appConfigLoader := AppConfigLoader{&fileOpsMock}

	fileOpsMock.On("ReadFile", "config.json").Return("asDasdfaF").Once()

	_, err := appConfigLoader.LoadAppConfig("config.json")

	fileOpsMock.AssertExpectations(test)

	assert.Error(test, err, "TestLoadAppConfigFailedDeserialization Failed")
}
