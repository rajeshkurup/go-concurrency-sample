package appstats

import (
	"appconfig"
	"math/rand"
	"time"
)

/**
 * @brief Container for holding host response
 */
type StatsDo struct {
	Application  *string  `json:"Application,omitempty"`
	Version      *string  `json:"Version,omitempty"`
	Uptime       *int64   `json:"Uptime,omitempty"`
	RequestCount *int64   `json:"Request_Count,omitempty"`
	ErrorCount   *int64   `json:"Error_Count,omitempty"`
	SuccessCount *int64   `json:"Success_Count,omitempty"`
	SuccssRate   *float64 `json:"Succss_Rate,omitempty"`
}

/**
 * @brief Constructor for StatsDo used only for testing purpose.
 *			Hardcoding API response for testing purpose.
 * @param config Application configuration initialized during startup
 * @param application Name of the application
 * @return A object of StatsDo filled with API response
 */
func MakeStatsDo(appConfig *appconfig.AppConfig, application string) StatsDo {
	versions := appConfig.ApiResponse.Versions
	randSrc := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSrc)
	uptime := randGen.Int63n(999999999999)
	requestCount := randGen.Int63n(999999999999)
	errorCount := randGen.Int63n(99999999)
	successCount := requestCount - errorCount
	version := randGen.Intn(10)

	return StatsDo{
		Application:  &application,
		Version:      &versions[version],
		Uptime:       &uptime,
		RequestCount: &requestCount,
		ErrorCount:   &errorCount,
		SuccessCount: &successCount,
	}
}
