package service

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"github.com/famesensor/playground-go-transaction-hexagonal/repository"
)

type CreateUserService interface {
	Create(ctx context.Context, req *model.CreateUserReq) error
}

type createUserService struct {
	userRepo    repository.UserRepository
	addressRepo repository.AddressRepository
	txManager   *manager.Manager
}

func NewCreateUserService(userRepo repository.UserRepository, addressRepo repository.AddressRepository, txManager *manager.Manager) CreateUserService {
	return &createUserService{
		userRepo,
		addressRepo,
		txManager,
	}
}

func (s *createUserService) Create(ctx context.Context, req *model.CreateUserReq) error {
	return s.txManager.Do(ctx, func(c context.Context) error {
		if err := s.userRepo.Create(c, &model.CreateUser{Name: req.Name}); err != nil {
			return err
		}

		// UserID set 1 for test error and rollback
		if err := s.addressRepo.Create(c, &model.CreateAddress{UserID: 0, Address: req.Address}); err != nil {
			return err
		}

		return nil
	})
}
