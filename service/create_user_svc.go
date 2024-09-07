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
	userRepo    repository.UserRepository
	addressRepo repository.AddressRepository
}

func NewCreateUserService(userRepo repository.UserRepository, addressRepo repository.AddressRepository) CreateUserService {
	return &createUserService{
		userRepo,
		addressRepo,
	}
}

func (s *createUserService) Create(ctx context.Context, req *model.CreateUserReq) error {

	tx := s.userRepo.Begin()

	if err := s.userRepo.WithTrx(tx).Create(ctx, &model.CreateUser{Name: req.Name}); err != nil {
		tx.Rollback()
		return err
	}

	// UserID set 1 for test error and rollback
	if err := s.addressRepo.WithTrx(tx).Create(ctx, &model.CreateAddress{UserID: 0, Address: req.Address}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
