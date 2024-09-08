package service

import (
	"context"

	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"github.com/famesensor/playground-go-transaction-hexagonal/repository"
)

type CreateUserService interface {
	Create(ctx context.Context, req *model.CreateUserReq) error
}

type createUserService struct {
	userRepo repository.UserRepository
}

func NewCreateUserService(userRepo repository.UserRepository) CreateUserService {
	return &createUserService{
		userRepo,
	}
}

func (s *createUserService) Create(ctx context.Context, req *model.CreateUserReq) error {
	return s.userRepo.CreateUserAndAddress(ctx, &model.CreateUser{Name: req.Name, Address: req.Address})
}
