package main

import (
	"log"

	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/server"
)

func main() {
	apiServer, err := server.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer apiServer.Shutdown()
	apiServer.Start(80)
}
