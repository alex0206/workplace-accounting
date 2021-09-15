package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/alex0206/workplace-accounting/internal/api"
	"github.com/alex0206/workplace-accounting/internal/services"
)

var serverHost = flag.String(
	"host",
	"https://workplace-accounting.herokuapp.com",
	"server host for updating workplace. Ex. ./client -host https://workplace.com",
)

func main() {
	flag.Parse()

	apiClient := api.NewWorkplaceAPIClient(*serverHost)
	clientServer := services.NewClientService(apiClient)
	wp, err := clientServer.UpdateWorkplace()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Workplace was successfully updated")

	wpBytes, err := json.MarshalIndent(wp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(wpBytes))
}
