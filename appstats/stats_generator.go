package appstats

import (
	"appconfig"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

/**
 * @brief Interface for Stats Generator
 */
type StatsGeneratorInterface interface {
	CollectStats(hostName string) []StatsDo
}

/**
 * @brief Makes API call to host and collect stats
 */
type StatsGenerator struct {
	appConfig *appconfig.AppConfig
	apiCaller ApiCallerInterface
}

/**
 * @brief Constructor for StatsGenerator
 * @param appConfig Application configuration initialized during startup
 * @param hostName Name of the host
 */
func MakeStatsGenerator(appConfig *appconfig.AppConfig) StatsGenerator {
	apiCaller := MakeApiCaller()
	return StatsGenerator{appConfig, &apiCaller}
}

/**
 * @brief A public method to make API call to host and collect stats.
 *			Deserialize stats received from host.
 * @param hostName Name of the host from which Stats to be collected
 * @return An array of StatsDo objects
 */
func (statsGenerator *StatsGenerator) CollectStats(hostName string) []StatsDo {
	hostStats := []StatsDo{}
	stats := ""
	url := fmt.Sprintf(statsGenerator.appConfig.DataSource.UrlFormat, hostName)
	for attempt := 0; len(stats) == 0 && attempt < statsGenerator.appConfig.DataSource.MaxAttempts; attempt++ {
		stats = statsGenerator.getStats(url)
		if len(stats) > 0 {
			err1 := json.Unmarshal([]byte(stats), &hostStats)
			if err1 != nil {
				log.Println(fmt.Sprintf("Failed to decode stats from %s - attempt=%d - error=%s", url, attempt, err1.Error()))
			} else {
				log.Println(fmt.Sprintf("Stats Collected from %s - attempt=%d", url, attempt))

				// Calculate Success Rate for all Applications
				for idx := 0; idx < len(hostStats); idx++ {
					appStats := &hostStats[idx]
					successRate := (float64(*appStats.SuccessCount) / float64(*appStats.RequestCount)) * 100.0
					appStats.SuccssRate = &successRate
				}
			}
		} else {
			log.Println(fmt.Sprintf("Failed to collect stats from %s - attempt=%d", url, attempt))
			// Wait before next attempt
			time.Sleep(time.Duration(statsGenerator.appConfig.DataSource.WaitTimeInSeconds))
		}
	}

	return hostStats
}

/**
 * @brief A private method to make API call to host and collect stats.
 * @param url API URL to collect Stats
 * @return Stats from host as string object
 */
func (statsGenerator *StatsGenerator) getStats(url string) string {
	stats := ""
	resp, err1 := statsGenerator.apiCaller.HttpGet(url)
	if err1 != nil {
		log.Println(fmt.Sprintf("Failed to connect to %s - error=%s", url, err1.Error()))
	} else {
		body, err2 := statsGenerator.apiCaller.IoRead(resp.Body)
		if err2 != nil {
			log.Println(fmt.Sprintf("Failed to read stats from %s - error=%s", url, err2.Error()))
		} else {
			stats = string(body)

			// Hardcoding API response for testing purpose
			// hostStats := []StatsDo{}
			// for _, appName := range statsGenerator.appConfig.ApiResponse.Applications {
			// 	hostStats = append(hostStats, MakeStatsDo(statsGenerator.appConfig, appName))
			// }

			// result, _ := json.Marshal(hostStats)
			// stats = string(result)
		}
	}

	return stats
}
