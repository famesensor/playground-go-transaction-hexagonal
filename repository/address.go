package repository

import (
	"context"

	trmgorm "github.com/avito-tech/go-transaction-manager/drivers/gorm/v2"
	"github.com/famesensor/playground-go-transaction-hexagonal/entity"
	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(ctx context.Context, req *model.CreateAddress) error
}

type addressRepository struct {
	db     *gorm.DB
	getter *trmgorm.CtxGetter
}

func NewAddress(db *gorm.DB, getter *trmgorm.CtxGetter) AddressRepository {
	return &addressRepository{db, getter}
}

func (a addressRepository) Create(ctx context.Context, req *model.CreateAddress) error {
	return a.getter.DefaultTrOrDB(ctx, a.db).WithContext(ctx).Create(&entity.Address{Address: &req.Address, UserID: int32(req.UserID)}).Error
}
