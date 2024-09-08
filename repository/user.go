package repository

import (
	"context"

	"github.com/famesensor/playground-go-transaction-hexagonal/entity"
	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, req *model.CreateUser) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) Create(ctx context.Context, req *model.CreateUser) error {
	return u.db.WithContext(ctx).Create(&entity.User{Name: req.Name}).Error
}
