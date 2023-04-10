package appconfig

/**
 * @brief Container to hold Execution section of configuration
 */
type Execution struct {
	ThreadPoolSize int `json:"ThreadPoolSize"`
}

/**
 * @brief Container to hold DataSource section of configuration
 */
type DataSource struct {
	MaxAttempts       int    `json:"MaxAttempts"`
	UrlFormat         string `json:"UrlFormat"`
	WaitTimeInSeconds int    `json:"WaitTimeInSeconds"`
}

/**
 * @brief Container to hold ApiResponse section of configuration
 */
type ApiResponse struct {
	Applications []string `json:"Applications"`
	Versions     []string `json:"Versions"`
}

/**
 * @brief Container to hold Application section of configuration
 */
type Application struct {
	HostFile   string `json:"HostFile"`
	ReportFile string `json:"ReportFile"`
}

/**
 * @brief Container to hold application configuration
 */
type AppConfig struct {
	Application Application `json:"Application"`
	ApiResponse ApiResponse `json:"ApiResponse"`
	DataSource  DataSource  `json:"DataSource"`
	Execution   Execution   `json:"Execution"`
}
