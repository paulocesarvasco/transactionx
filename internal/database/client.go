package database

import (
	"errors"
	"log"
	"transactionx/internal/resources"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client interface {
	RegisterTransaction(t resources.Transaction) (resources.Transaction, error)
	RetrieveTransactions() ([]resources.Transaction, error)
	SearchTransaction(id string) (resources.Transaction, error)
}

type dbClient struct {
	db *gorm.DB
}

func NewSQLiteClient() Client {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	db.AutoMigrate(&resources.Transaction{})

	return &dbClient{db: db}
}

func (c *dbClient) RegisterTransaction(t resources.Transaction) (resources.Transaction, error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return resources.Transaction{}, tx.Error
	}

	if err := tx.Create(&t).Error; err != nil {
		tx.Rollback()
		return resources.Transaction{}, err
	}
	err := tx.Commit().Error
	if err != nil {
		return resources.Transaction{}, err
	}
	return t, nil
}

func (c dbClient) RetrieveTransactions() ([]resources.Transaction, error) {
	var t []resources.Transaction
	result := c.db.Find(&t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

func (c dbClient) SearchTransaction(id string) (resources.Transaction, error) {
	var t resources.Transaction
	result := c.db.Where("id = ?", id).First(&t)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return resources.Transaction{}, nil
	}
	if result.Error != nil {
		return resources.Transaction{}, result.Error
	}
	return t, nil

}
