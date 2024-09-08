package repository

import (
	"context"

	trmgorm "github.com/avito-tech/go-transaction-manager/drivers/gorm/v2"
	"github.com/famesensor/playground-go-transaction-hexagonal/entity"
	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, req *model.CreateUser) error
}

type userRepository struct {
	db     *gorm.DB
	getter *trmgorm.CtxGetter
}

func NewUser(db *gorm.DB, getter *trmgorm.CtxGetter) UserRepository {
	return &userRepository{db, getter}
}

func (u userRepository) Create(ctx context.Context, req *model.CreateUser) error {
	return u.getter.DefaultTrOrDB(ctx, u.db).WithContext(ctx).Create(&entity.User{Name: req.Name}).Error
}
