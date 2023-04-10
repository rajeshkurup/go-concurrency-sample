package main

import (
	"appmain"
	"log"
)

/**
 * @brief Entry point for stats_generator application.
 */
func main() {
	log.Println("Stats Generator")

	appMain := appmain.MakeAppMain()
	appMain.Run()

	log.Println("All Done")
}
