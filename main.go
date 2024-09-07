package main

import (
	"fmt"
	"log"
	"net"

	"github.com/famesensor/playground-go-transaction-hexagonal/database"
	"github.com/famesensor/playground-go-transaction-hexagonal/handler"
	"github.com/famesensor/playground-go-transaction-hexagonal/proto"
	"github.com/famesensor/playground-go-transaction-hexagonal/repository"
	"github.com/famesensor/playground-go-transaction-hexagonal/service"
	"google.golang.org/grpc"
)

func main() {

	db, err := database.InitPostgres()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUser(db)
	addressRepo := repository.NewAddress(db)

	createUserSvc := service.NewCreateUserService(userRepo, addressRepo)

	port := 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc.NewServer()
	s := grpc.NewServer()
	proto.RegisterUserServer(s, handler.NewUserHandler(createUserSvc))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
