package main

import (
	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/server"
)

func main() {
	apiServer := server.Init()
	apiServer.Start(80)
}
