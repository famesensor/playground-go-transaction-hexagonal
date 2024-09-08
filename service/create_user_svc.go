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
	uowRepo     repository.UnitOfWork
	userRepo    repository.UserRepository
	addressRepo repository.AddressRepository
}

func NewCreateUserService(uowRepo repository.UnitOfWork, userRepo repository.UserRepository, addressRepo repository.AddressRepository) CreateUserService {
	return &createUserService{
		uowRepo,
		userRepo,
		addressRepo,
	}
}

func (s *createUserService) Create(ctx context.Context, req *model.CreateUserReq) error {
	return s.uowRepo.BeginTx(ctx, func(tx repository.UnitOfWork) error {
		if err := tx.UserRepository().Create(ctx, &model.CreateUser{Name: req.Name}); err != nil {
			return err
		}
		if err := tx.AddressRepository().Create(ctx, &model.CreateAddress{UserID: 0, Address: req.Address}); err != nil {
			return err
		}
		return nil
	})
}
