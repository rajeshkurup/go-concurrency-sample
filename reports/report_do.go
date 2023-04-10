package reports

import (
	"appstats"
	"fmt"
)

/**
 * @brief Container to hold aggregaed report
 */
type ReportDo struct {
	report map[string]map[string]SuccessRateDo
}

/**
 * @brief Constructor for ReportDo
 * @return An object of ReportDo
 */
func MakeReportDo() ReportDo {
	return ReportDo{make(map[string]map[string]SuccessRateDo)}
}

/**
 * @brief A Public method to get actual Success Rateb report from ReportDo
 * @return Map of Average Success Rate based on Application and Version
 */
func (reportDo *ReportDo) GetReport() map[string]map[string]SuccessRateDo {
	return reportDo.report
}

/**
 * @brief A public method to add stats from an application running on a host into report for aggregation
 * @param stats Stats from an application running on a host
 * @return Pointer to current ReportDo object
 */
func (reportDo *ReportDo) AddStats(stats appstats.StatsDo) *ReportDo {
	if appStats, appOk := reportDo.report[*stats.Application]; appOk {
		if versionStats, versionOk := appStats[*stats.Version]; versionOk {
			reportDo.report[*stats.Application][*stats.Version] = versionStats.AddSuccessRate(*stats.SuccssRate)
		} else {
			reportDo.report[*stats.Application][*stats.Version] = MakeSuccessRateDo(*stats.SuccssRate, 1.0)
		}
	} else {
		reportDo.report[*stats.Application] = make(map[string]SuccessRateDo)
		reportDo.report[*stats.Application][*stats.Version] = MakeSuccessRateDo(*stats.SuccssRate, 1.0)
	}

	return reportDo
}

/**
 * @brief A public method to build actual report from aggregated stats
 * @return Pointer to current ReportDo object
 */
func (reportDo *ReportDo) BuildReport() *ReportDo {
	for appName, appStats := range reportDo.report {
		for version := range appStats {
			versionStats := reportDo.report[appName][version]
			reportDo.report[appName][version] = *versionStats.Finalize()
		}
	}

	return reportDo
}

/**
 * @brief A public method to print report in human readable form on console
 */
func (reportDo *ReportDo) PrettyPrint() {
	fmt.Println("============================================")
	fmt.Println("********** Success Rate Report *************")
	fmt.Println("============================================")
	fmt.Println("| Application | Version | Success Rate (%) |")
	fmt.Println("============================================")

	for appName, appStats := range reportDo.report {
		for version, versionStats := range appStats {
			fmt.Println(fmt.Sprintf("| %-11s | %-7s | %16.5f |", appName, version, *versionStats.AverageSuccessRate))
		}
	}

	fmt.Println("============================================")
}
