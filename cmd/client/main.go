package main

import (
	"flag"
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
	if err := clientServer.UpdateWorkplace(); err != nil {
		log.Fatal(err)
	}
}
