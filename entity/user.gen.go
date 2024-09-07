// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string     `gorm:"column:name;not null" json:"name"`
	CreatedDate time.Time  `gorm:"column:created_date;not null;default:now()" json:"created_date"`
	UpdatedDate *time.Time `gorm:"column:updated_date" json:"updated_date"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
