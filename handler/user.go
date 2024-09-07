package handler

import (
	"context"
	"time"

	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"github.com/famesensor/playground-go-transaction-hexagonal/proto"
	"github.com/famesensor/playground-go-transaction-hexagonal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userHandler struct {
	createUserSvc service.CreateUserService
	proto.UnimplementedUserServer
}

func NewUserHandler(createUserSvc service.CreateUserService) proto.UserServer {
	return &userHandler{createUserSvc: createUserSvc}
}

func (u *userHandler) CreateUserHandler(ctx context.Context, req *proto.CreateUserReq) (*proto.CreateUserRes, error) {
	if err := u.createUserSvc.Create(ctx, &model.CreateUserReq{Name: req.Name, Address: req.Address}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateUserRes{
		Data: &proto.Timestamp{
			Timestamp: time.Now().String(),
		},
	}, nil
}
