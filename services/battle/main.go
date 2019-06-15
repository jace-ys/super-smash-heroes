package main

import (
	"github.com/jace-ys/super-smash-heroes/libraries/go/service"
	"github.com/jace-ys/super-smash-heroes/services/battle/pkg/srvc"
)

func main() {
	s := srvc.NewService()
	service.StartServer(s, 80)
}
