package service

import (
	"fmt"
	"log"
	"net"
)

type Service interface {
	Init() error
	Serve(lis net.Listener) error
	Shutdown()
}

func StartServer(s Service, port int) {
	err := s.Init()
	if err != nil {
		log.Fatal(err)
	}
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
