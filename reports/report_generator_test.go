package reports

import (
	"appconfig"
	"appstats"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateReportSuccess(test *testing.T) {
	appConfig := appconfig.AppConfig{}
	appConfig.Execution.ThreadPoolSize = 2
	appConfig.ApiResponse.Applications = []string{"WebApp", "Cahce", "Database"}
	appConfig.ApiResponse.Versions = []string{"1.0.1", "2.2.2", "3.1.5"}

	statsGeneratorMock := appstats.StatsGeneratorMock{}
	reportGenerator := ReportGenerator{&appConfig, &statsGeneratorMock}

	successRate1 := float64(90.99)
	successRate2 := float64(91.92)
	successRate3 := float64(95.78)

	hosts := []string{"HostA", "HostB", "HostC"}

	stats1 := appstats.StatsDo{}
	stats1.SuccssRate = &successRate1
	stats1.Application = &appConfig.ApiResponse.Applications[0]
	stats1.Version = &appConfig.ApiResponse.Versions[0]
	statsColl1 := []appstats.StatsDo{stats1}

	stats2 := appstats.StatsDo{}
	stats2.SuccssRate = &successRate2
	stats2.Application = &appConfig.ApiResponse.Applications[1]
	stats2.Version = &appConfig.ApiResponse.Versions[1]
	statsColl2 := []appstats.StatsDo{stats2}

	stats3 := appstats.StatsDo{}
	stats3.SuccssRate = &successRate3
	stats3.Application = &appConfig.ApiResponse.Applications[2]
	stats3.Version = &appConfig.ApiResponse.Versions[2]
	statsColl3 := []appstats.StatsDo{stats3}

	statsGeneratorMock.On("CollectStats", "HostA").Return(statsColl1).Once()
	statsGeneratorMock.On("CollectStats", "HostB").Return(statsColl2).Once()
	statsGeneratorMock.On("CollectStats", "HostC").Return(statsColl3).Once()

	reportResult := reportGenerator.GenerateReport(hosts)

	statsGeneratorMock.AssertExpectations(test)

	assert.Equal(test, successRate1, *reportResult.report["WebApp"]["1.0.1"].AverageSuccessRate, "TestGenerateReportSuccess Failed: Wrong AverageSuccessRate for WebApp")
	assert.Equal(test, successRate2, *reportResult.report["Cahce"]["2.2.2"].AverageSuccessRate, "TestGenerateReportSuccess Failed: Wrong AverageSuccessRate for Cahce")
	assert.Equal(test, successRate3, *reportResult.report["Database"]["3.1.5"].AverageSuccessRate, "TestGenerateReportSuccess Failed: Wrong AverageSuccessRate for Database")
}

func TestGenerateReportFailed(test *testing.T) {
	appConfig := appconfig.AppConfig{}
	appConfig.Execution.ThreadPoolSize = 2
	appConfig.ApiResponse.Applications = []string{"WebApp", "Cahce", "Database"}
	appConfig.ApiResponse.Versions = []string{"1.0.1", "2.2.2", "3.1.5"}

	statsGeneratorMock := appstats.StatsGeneratorMock{}
	reportGenerator := ReportGenerator{&appConfig, &statsGeneratorMock}

	hosts := []string{"HostA", "HostB", "HostC"}

	statsGeneratorMock.On("CollectStats", "HostA").Return([]appstats.StatsDo{}).Once()
	statsGeneratorMock.On("CollectStats", "HostB").Return([]appstats.StatsDo{}).Once()
	statsGeneratorMock.On("CollectStats", "HostC").Return([]appstats.StatsDo{}).Once()

	reportResult := reportGenerator.GenerateReport(hosts)

	statsGeneratorMock.AssertExpectations(test)

	assert.Equal(test, 0, len(reportResult.report), "TestGenerateReportFailed Failed: Stats found")
}
