# Stats Report Generator

## Functions
* Reads list of hosts from `hosts.txt` under root folder
* Makes API calls to all hosts in the list
* Retries when there is a failure in API call
* Collects stats from all hosts in the list
* Aggregates stats from all hosts in the list
* Generates stats report as json file `stats.json` under root folder
* Prints stats report in humal readable form on console

## Assumptions
* API response can be hardcoded as there no real hosts existing

## Configurations
* Appication configurations as read from `config.json` and deserialized into `appconfig.AppConfig`

### Application
* `HostFile`: Fully qualified path to `hosts.txt`
* `ReportFile`: Fully qualified path to `stats.json`

### ApiResponse
* `Applications`: Hardcoded list of applications as there are no real hosts existing
* `Versions`: Hardcoded list of versions as there are no real hosts existing

### DataSource
* `MaxAttempts`: Maximum number of attempts would be made by the application to obtain stats from a host
* `UrlFormat`: URL of the Data Source without Host name
* `WaitTimeInSeconds`: Time in seconds to wait before making next attempt on host in case of a failure

### Execution
* `ThreadPoolSize`: Maximum number of threads would be spawned by the application. Limits depends on CPU and Memory capacity of system where the process runs.

## Entry Point
* Refer to `main` function in `main.go` under root folder

## Packages and Classes

### appmain
* `AppMain`: Stats Generator application controller

### appconfig
* `AppConfig`: Data Object to hold application configurations
* `AppConfigLoader`: Loads application configurations from `config.json`

### appstats
* `StatsGenerator`: Make API calls and read stats from hosts
* `StatsDo`: Data Object to hold API response
* `ApiCaller`: Adaptor to handle HTTP operations

### reports
* `ReportGenerator`: Generate report from list of `appstats.StatsDo`
* `ReportDo`: Data Object to hold report
* `SuccessRateDo`: Data Object to hold aggregated Success Rate

### fileops
* `FileOps`: File operation helper to read and write files
* `OsFileOps`: File operation adaptor to handle `os` file operations

## Build
* Run `make build` from root folder

## Test
* Run `make test` from root folder

## Run
* Run `make run` from root folder

## Clean
* Run `make clean` from root folder

## Requirements
* GoLang version `1.20.3`
* Configurations written as json in `config.json` present in root folder
* List of hosts written as text file in `hosts.txt` present in root folder
