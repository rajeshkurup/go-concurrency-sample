package reports

import (
	"appconfig"
	"appstats"
	"log"
	"math"
	"sync"
)

/**
 * @brief Interface for Report Generator
 */
type ReportGeneratorInterface interface {
	GenerateReport(hosts []string) ReportDo
}

/**
 * @brief Generates Success Rate Report
 */
type ReportGenerator struct {
	appConfig      *appconfig.AppConfig
	statsGenerator appstats.StatsGeneratorInterface
}

/**
 * @brief Constructor for ReportGenerator
 * @param appConfig Application configuration initialized during startup
 * @return An object of ReportGenerator
 */
func MakeReportGenerator(appConfig *appconfig.AppConfig) ReportGenerator {
	statsGenerator := appstats.MakeStatsGenerator(appConfig)
	return ReportGenerator{appConfig, &statsGenerator}
}

/**
 * @brief A private method to collect stats from hosts.
 * 			Creates thread pool and calls hosts in parallel.
 *			Maximum number of threads would be taken from configuration.
 * 			Consolidates stats from different hosts into one list of appstats.StatsDo objects.
 * @param hosts List of hosts
 * @return A list of appstats.StatsDo objects containing reponses from hosts
 */
func (reportGenerator *ReportGenerator) collectStats(hosts []string) []appstats.StatsDo {
	log.Println("Start Collecting Stats")
	hostSlices := int(len(hosts) / reportGenerator.appConfig.Execution.ThreadPoolSize)
	reminder := math.Mod(float64(len(hosts)), float64(reportGenerator.appConfig.Execution.ThreadPoolSize))
	if int(reminder) > 0 {
		hostSlices++
	}

	statsCollection := []appstats.StatsDo{}
	for i := 0; i < hostSlices; i++ {
		startIdx := i * reportGenerator.appConfig.Execution.ThreadPoolSize
		waitGroup := sync.WaitGroup{}
		futureStats := make(chan []appstats.StatsDo)
		for j := startIdx; j < startIdx+reportGenerator.appConfig.Execution.ThreadPoolSize && j < len(hosts); j++ {
			if len(hosts[j]) > 0 {
				waitGroup.Add(1)
				go reportGenerator.collectStatsAsync(hosts[j], &waitGroup, futureStats)
			}
		}

		go func() {
			log.Println("Waiting to Collect Stats")
			waitGroup.Wait()
			close(futureStats)
		}()

		for stats := range futureStats {
			statsCollection = append(statsCollection, stats...)
		}
	}

	log.Println("Stats Collection Done")
	return statsCollection
}

/**
 * @brief A private method to make call to given host asynchronously.
 * @param host Name of the host
 * @param waitGroup The thread pool created in collectStats method
 * @param stats A channel (promise or future) to receive stats from host when done
 */
func (reportGenerator *ReportGenerator) collectStatsAsync(host string, waitGroup *sync.WaitGroup, stats chan []appstats.StatsDo) {
	defer waitGroup.Done()

	stats <- reportGenerator.statsGenerator.CollectStats(host)
}

/**
 * @brief A public method to collect stats from hosts and generate report from them.
 * @param hosts List of hosts
 * @return An object of ReportDo containing aggregated report.
 */
func (reportGenerator *ReportGenerator) GenerateReport(hosts []string) ReportDo {
	hostStats := reportGenerator.collectStats(hosts)

	report := MakeReportDo()
	for _, appStats := range hostStats {
		report.AddStats(appStats)
	}

	return *report.BuildReport()
}
