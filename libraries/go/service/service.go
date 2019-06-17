package service

import (
	"fmt"
	"log"
	"net"
)

type Service interface {
	Init()
	Serve(lis net.Listener) error
	Shutdown()
}

func StartServer(s Service, port int) {
	s.Init()
	defer s.Shutdown()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listening on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
