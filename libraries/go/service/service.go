package service

import (
	"fmt"
	"log"
	"net"
)

type Service interface {
	Register()
	Serve(lis net.Listener) error
}

func StartServer(s Service, port int) {
	s.Register()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Service listening on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
