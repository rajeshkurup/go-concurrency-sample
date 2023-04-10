package appstats

import (
	"appconfig"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCollectStatsSuccess(test *testing.T) {
	appConfig := appconfig.AppConfig{}
	appConfig.DataSource.UrlFormat = "http://%s/stats"
	appConfig.DataSource.MaxAttempts = 2
	appConfig.DataSource.WaitTimeInSeconds = 1
	appConfig.ApiResponse.Applications = []string{"WebApp", "Cahce"}
	appConfig.ApiResponse.Versions = []string{"1.0.1", "2.2.2"}

	apiCallerMock := ApiCallerMock{}
	statsGenerator := StatsGenerator{&appConfig, &apiCallerMock}

	errorCount := int64(10)
	successCount := int64(90)
	requestCount := errorCount + successCount
	uptime := int64(1000)
	successRate := (float64(successCount) / float64(requestCount)) * 100.0

	httpResp := http.Response{}

	stats := StatsDo{}
	stats.Application = &appConfig.ApiResponse.Applications[0]
	stats.ErrorCount = &errorCount
	stats.RequestCount = &requestCount
	stats.SuccessCount = &successCount
	stats.Version = &appConfig.ApiResponse.Versions[0]
	stats.Uptime = &uptime

	hostStats := []StatsDo{stats}
	hostStatsJson, _ := json.Marshal(hostStats)

	apiCallerMock.On("HttpGet", "http://HostA/stats").Return(&httpResp, nil).Once()
	apiCallerMock.On("IoRead", mock.Anything).Return(hostStatsJson, nil).Once()

	statsResult := statsGenerator.CollectStats("HostA")

	apiCallerMock.AssertExpectations(test)

	assert.Equal(test, appConfig.ApiResponse.Applications[0], *statsResult[0].Application, "TestCollectStatsSuccess Failed: Wrong Application")
	assert.Equal(test, appConfig.ApiResponse.Versions[0], *statsResult[0].Version, "TestCollectStatsSuccess Failed: Wrong Version")
	assert.Equal(test, errorCount, *statsResult[0].ErrorCount, "TestCollectStatsSuccess Failed: Wrong ErrorCount")
	assert.Equal(test, successCount, *statsResult[0].SuccessCount, "TestCollectStatsSuccess Failed: Wrong SuccessCount")
	assert.Equal(test, requestCount, *statsResult[0].RequestCount, "TestCollectStatsSuccess Failed: Wrong RequestCount")
	assert.Equal(test, uptime, *statsResult[0].Uptime, "TestCollectStatsSuccess Failed: Wrong Uptime")
	assert.Equal(test, successRate, *statsResult[0].SuccssRate, "TestCollectStatsSuccess Failed: Wrong SuccssRate")
}

func TestCollectStatsFailedIoRead(test *testing.T) {
	appConfig := appconfig.AppConfig{}
	appConfig.DataSource.UrlFormat = "http://%s/stats"
	appConfig.DataSource.MaxAttempts = 2
	appConfig.DataSource.WaitTimeInSeconds = 1
	appConfig.ApiResponse.Applications = []string{"WebApp", "Cahce"}
	appConfig.ApiResponse.Versions = []string{"1.0.1", "2.2.2"}

	apiCallerMock := ApiCallerMock{}
	statsGenerator := StatsGenerator{&appConfig, &apiCallerMock}

	httpResp := http.Response{}

	apiCallerMock.On("HttpGet", "http://HostA/stats").Return(&httpResp, nil).Twice()
	apiCallerMock.On("IoRead", mock.Anything).Return([]byte(""), errors.New("Failed to read API Response")).Twice()

	statsResult := statsGenerator.CollectStats("HostA")

	apiCallerMock.AssertExpectations(test)

	assert.Equal(test, 0, len(statsResult), "TestCollectStatsFailedIoRead Failed: Wrong Response")
}

func TestCollectStatsFailedHttpGet(test *testing.T) {
	appConfig := appconfig.AppConfig{}
	appConfig.DataSource.UrlFormat = "http://%s/stats"
	appConfig.DataSource.MaxAttempts = 2
	appConfig.DataSource.WaitTimeInSeconds = 1
	appConfig.ApiResponse.Applications = []string{"WebApp", "Cahce"}
	appConfig.ApiResponse.Versions = []string{"1.0.1", "2.2.2"}

	apiCallerMock := ApiCallerMock{}
	statsGenerator := StatsGenerator{&appConfig, &apiCallerMock}

	httpResp := http.Response{}

	apiCallerMock.On("HttpGet", "http://HostA/stats").Return(&httpResp, errors.New("Http Get Failed")).Twice()

	statsResult := statsGenerator.CollectStats("HostA")

	apiCallerMock.AssertExpectations(test)

	assert.Equal(test, 0, len(statsResult), "TestCollectStatsFailedHttpGet Failed: Wrong Response")
}

func TestCollectStatsFailedDeserialization(test *testing.T) {
	appConfig := appconfig.AppConfig{}
	appConfig.DataSource.UrlFormat = "http://%s/stats"
	appConfig.DataSource.MaxAttempts = 2
	appConfig.DataSource.WaitTimeInSeconds = 1
	appConfig.ApiResponse.Applications = []string{"WebApp", "Cahce"}
	appConfig.ApiResponse.Versions = []string{"1.0.1", "2.2.2"}

	apiCallerMock := ApiCallerMock{}
	statsGenerator := StatsGenerator{&appConfig, &apiCallerMock}

	httpResp := http.Response{}

	apiCallerMock.On("HttpGet", "http://HostA/stats").Return(&httpResp, nil).Once()
	apiCallerMock.On("IoRead", mock.Anything).Return([]byte("adasfsaf"), nil).Once()

	statsResult := statsGenerator.CollectStats("HostA")

	apiCallerMock.AssertExpectations(test)

	assert.Equal(test, 0, len(statsResult), "TestCollectStatsFailedDeserialization Failed: Deserialization")
}
