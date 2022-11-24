package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test-golang/domain"
)

func InitGorm() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=disable",
		"localhost", "5432", "postgres", "test-golang", "AdminPassword123")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.Product{})
	if err != nil {
		panic(err)
	}

	return db
}
