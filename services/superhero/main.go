package main

import (
	"github.com/jace-ys/super-smash-heroes/libraries/go/config"
	"github.com/jace-ys/super-smash-heroes/libraries/go/service"
	"github.com/jace-ys/super-smash-heroes/services/superhero/pkg/srvc"
)

func main() {
	s := srvc.NewService()
	port := config.Get("service.superhero.port").Int(80)
	service.StartServer(s, port)
}
