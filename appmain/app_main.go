package appmain

import (
	"appconfig"
	"encoding/json"
	"fileops"
	"fmt"
	"log"
	"reports"
	"strings"
)

/**
 * @brief Controller for Stats Generator Application
 */
type AppMain struct {
	appConfigLoader appconfig.AppConfigLoaderInterface
	fileOps         fileops.FileOpsInterface
	reportGenerator reports.ReportGeneratorInterface
	appConfig       appconfig.AppConfig
}

/**
 * @brief Constructor for Main.
 *			Load Application Configuration from config.json.
 *			Stop execution if load config fails.
 * @return Instance of AppMain
 */
func MakeAppMain() AppMain {
	appConfigLoader := appconfig.MakeAppConfigLoader()
	appConfig, err := appConfigLoader.LoadAppConfig("./config.json")
	if err != nil {
		// Failed to load config. Log Error and Exit.
		log.Fatalln(fmt.Sprintf("Failed to read application configuration from ./config.json - error=%s", err.Error()))
	}

	fileOps := fileops.MakeFileOps()
	reportGenerator := reports.MakeReportGenerator(&appConfig)
	return AppMain{&appConfigLoader, &fileOps, &reportGenerator, appConfig}
}

/**
 * @brief Start up method of stats_Generator application.
 *			Create AppConfig object with configurations from config.json file.
 *			Load list of hosts from hosts.txt file.
 *			Generate Success Rate report.
 *			Save report as report.json file.
 *			Print report on console in human readable format.
 */
func (appMain *AppMain) Run() {
	hostList := appMain.getHosts(appMain.appConfig)
	if len(hostList) > 0 {
		log.Println("Start Generating Success Rate Report")

		report := appMain.reportGenerator.GenerateReport(hostList)

		jsonReport, _ := json.Marshal(report.GetReport())
		appMain.fileOps.WriteFile(appMain.appConfig.Application.ReportFile, jsonReport)
		report.PrettyPrint()
	} else {
		log.Println(fmt.Sprintf("Failed to read list of hosts from %s", appMain.appConfig.Application.HostFile))
	}
}

/**
 * @brief Helper function to read hosts.txt and build list of hosts
 * @param appConfig Instance of Application Configuration
 * @return Array of string objects containing host names
 */
func (appMain *AppMain) getHosts(appConfig appconfig.AppConfig) []string {
	log.Println("Reading host list")

	hostList := []string{}
	hosts := appMain.fileOps.ReadFile(appConfig.Application.HostFile)
	if len(hosts) > 0 {
		hostList = strings.Split(hosts, "\n")
	}

	return hostList
}
