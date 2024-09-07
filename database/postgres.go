package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgres() (*gorm.DB, error) {
	dsn := "host=localhost user=username password=pass dbname=psql_db port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		CreateBatchSize:        100,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
