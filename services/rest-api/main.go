package main

import (
	"log"

	"github.com/jace-ys/super-smash-heroes/libraries/go/config"
	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/server"
)

func main() {
	apiServer, err := server.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer apiServer.Shutdown()
	port := config.Get("service.rest-api.port").Int(80)
	apiServer.Start(port)
}
