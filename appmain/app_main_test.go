package appmain

import (
	"appconfig"
	"fileops"
	"reports"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestRunSuccess(test *testing.T) {
	appConfig := appconfig.AppConfig{}
	appConfig.Application.HostFile = "hosts.txt"
	appConfig.Application.ReportFile = "report.json"
	appConfig.ApiResponse.Applications = []string{"WebApp", "Cahce", "Database"}
	appConfig.ApiResponse.Versions = []string{"1.0.1", "2.2.2", "3.1.5"}

	reportGeneratorMock := reports.ReportGeneratorMock{}
	fileOpsMock := fileops.FileOpsMock{}
	appConfigLoaderMock := appconfig.AppConfigLoaderMock{}

	appMain := AppMain{&appConfigLoaderMock, &fileOpsMock, &reportGeneratorMock, appConfig}

	hosts := []string{"HostA", "HostB"}

	report := reports.MakeReportDo()

	fileOpsMock.On("ReadFile", "hosts.txt").Return("HostA\nHostB").Once()
	fileOpsMock.On("WriteFile", "report.json", mock.Anything).Return(nil).Once()
	reportGeneratorMock.On("GenerateReport", hosts).Return(report).Once()

	appMain.Run()

	reportGeneratorMock.AssertExpectations(test)
	fileOpsMock.AssertExpectations(test)
}

func TestRunFailed(test *testing.T) {
	appConfig := appconfig.AppConfig{}
	appConfig.Application.HostFile = "hosts.txt"
	appConfig.Application.ReportFile = "report.json"
	appConfig.ApiResponse.Applications = []string{"WebApp", "Cahce", "Database"}
	appConfig.ApiResponse.Versions = []string{"1.0.1", "2.2.2", "3.1.5"}

	reportGeneratorMock := reports.ReportGeneratorMock{}
	fileOpsMock := fileops.FileOpsMock{}
	appConfigLoaderMock := appconfig.AppConfigLoaderMock{}

	appMain := AppMain{&appConfigLoaderMock, &fileOpsMock, &reportGeneratorMock, appConfig}

	fileOpsMock.On("ReadFile", "hosts.txt").Return("").Once()

	appMain.Run()

	fileOpsMock.AssertExpectations(test)
}
