package repository

import "gorm.io/gorm"

type Transactor interface {
	Begin() *gorm.DB
}
