package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UnitOfWork interface {
	BeginTx(ctx context.Context, fn func(tx UnitOfWork) error) error
	UserRepository() UserRepository
	AddressRepository() AddressRepository
	withTx(*gorm.DB) UnitOfWork
}

type unitOfWork struct {
	db          *gorm.DB
	userRepo    UserRepository
	addressRepo AddressRepository
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{
		db:          db,
		userRepo:    NewUser(db),
		addressRepo: NewAddress(db),
	}
}

// BeginTx runs a function within a transaction. If the function returns an
// error, the transaction is rolled back. If the function does not return an
// error, the transaction is committed.
func (u *unitOfWork) BeginTx(ctx context.Context, fn func(tx UnitOfWork) error) error {
	var err error
	tx := u.db.Begin()

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback().Error; rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit().Error
		}
	}()

	err = fn(u.withTx(tx))
	return err
}

// AddressRepository returns the AddressRepository instance.
func (u *unitOfWork) AddressRepository() AddressRepository {
	return u.addressRepo
}

// UserRepository returns the UserRepository instance.
func (u *unitOfWork) UserRepository() UserRepository {
	return u.userRepo
}

func (u *unitOfWork) withTx(tx *gorm.DB) UnitOfWork {
	return &unitOfWork{
		db:          tx,
		userRepo:    NewUser(tx),
		addressRepo: NewAddress(tx),
	}
}
