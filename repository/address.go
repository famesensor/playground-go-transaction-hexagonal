package repository

import (
	"context"

	"github.com/famesensor/playground-go-transaction-hexagonal/entity"
	"github.com/famesensor/playground-go-transaction-hexagonal/model"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Transactor
	WithTrx(tx *gorm.DB) AddressRepository
	Create(ctx context.Context, req *model.CreateAddress) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddress(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (a addressRepository) Begin() *gorm.DB {
	return a.db.Begin()
}

func (a addressRepository) WithTrx(tx *gorm.DB) AddressRepository {
	if tx != nil {
		a.db = tx
	}

	return a
}

func (a addressRepository) Create(ctx context.Context, req *model.CreateAddress) error {
	return a.db.WithContext(ctx).Create(&entity.Address{Address: &req.Address, UserID: int32(req.UserID)}).Error
}
