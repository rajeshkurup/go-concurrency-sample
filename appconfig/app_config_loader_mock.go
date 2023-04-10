package appconfig

import (
	"github.com/stretchr/testify/mock"
)

type AppConfigLoaderMock struct {
	mock.Mock
}

func (appConfigLoaderMock *AppConfigLoaderMock) LoadAppConfig(configFilePath string) (AppConfig, error) {
	args := appConfigLoaderMock.Called(configFilePath)
	return args.Get(0).(AppConfig), args.Error(1)
}
