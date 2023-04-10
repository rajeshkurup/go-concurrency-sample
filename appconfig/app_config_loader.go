package appconfig

import (
	"encoding/json"
	"errors"
	"fileops"
	"fmt"
	"log"
)

/**
 * @brief Interface for Application Configuration Loader
 */
type AppConfigLoaderInterface interface {
	LoadAppConfig(configFilePath string) (AppConfig, error)
}

/**
 * @brief Helps to load Application Configuration
 */
type AppConfigLoader struct {
	fileOps fileops.FileOpsInterface
}

/**
 * @brief Constructor for AppConfigLoader
 * @return An object of AppConfigLoader
 */
func MakeAppConfigLoader() AppConfigLoader {
	fileOps := fileops.MakeFileOps()
	return AppConfigLoader{&fileOps}
}

/**
 * @brief Load Application Configuration from given JSON file
 * @param configFilePath Fully qualified path to configuration JSON file
 * @return Instance of AppConfig if succeeded
 * @return Instance of error if failed
 */
func (appConfigLoader *AppConfigLoader) LoadAppConfig(configFilePath string) (AppConfig, error) {
	log.Println(fmt.Sprintf("Loading Application Configuration from %s", configFilePath))
	appConfig := AppConfig{}
	err := errors.New("")

	appConfigJson := appConfigLoader.fileOps.ReadFile(configFilePath)
	if len(appConfigJson) > 0 {
		err = json.Unmarshal([]byte(appConfigJson), &appConfig)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to deserialize application configuaration %s - error=%s", appConfigJson, err.Error()))
		}
	} else {
		errMsg := fmt.Sprintf("Failed to read application configuaration file %s - error=%s", configFilePath, err.Error())
		log.Println(errMsg)
		err = errors.New(errMsg)
	}

	return appConfig, err
}
