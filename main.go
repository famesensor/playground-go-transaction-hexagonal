package main

import (
	"fmt"
	"log"
	"net"

	trmgorm "github.com/avito-tech/go-transaction-manager/drivers/gorm/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/avito-tech/go-transaction-manager/trm/v2/settings"
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

	userRepo := repository.NewUser(db, trmgorm.DefaultCtxGetter)
	addressRepo := repository.NewAddress(db, trmgorm.DefaultCtxGetter)

	trManager := manager.Must(
		trmgorm.NewDefaultFactory(db),
		manager.WithSettings(trmgorm.MustSettings(
			settings.Must(
				settings.WithPropagation(trm.PropagationNested))),
		),
	)

	createUserSvc := service.NewCreateUserService(userRepo, addressRepo, trManager)

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
