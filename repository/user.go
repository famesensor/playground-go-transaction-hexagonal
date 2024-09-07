package repository

import (
	"context"

	"github.com/famesensor/playground-go-transaction-hexagonal/entity"
	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Transactor
	WithTrx(tx *gorm.DB) UserRepository
	Create(ctx context.Context, req *model.CreateUser) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) Begin() *gorm.DB {
	return u.db.Begin()
}

func (u userRepository) WithTrx(tx *gorm.DB) UserRepository {
	if tx != nil {
		u.db = tx
	}

	return u
}

func (u userRepository) Create(ctx context.Context, req *model.CreateUser) error {
	return u.db.WithContext(ctx).Create(&entity.User{Name: req.Name}).Error
}
