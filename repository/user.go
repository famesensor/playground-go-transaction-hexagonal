package repository

import (
	"context"

	"github.com/famesensor/playground-go-transaction-hexagonal/entity"
	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserAndAddress(ctx context.Context, req *model.CreateUser) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUserAndAddress(ctx context.Context, req *model.CreateUser) error {
	tx := u.db.WithContext(ctx).Begin()

	var done bool

	defer func() {
		if !done {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&entity.User{Name: req.Name}).Error; err != nil {
		// tx.Rollback()
		return err
	}

	if err := tx.Create(&entity.Address{UserID: 0, Address: &req.Address}).Error; err != nil {
		// tx.Rollback()
		return err
	}

	done = true
	return tx.Commit().Error
}
