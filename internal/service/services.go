package service

import (
	"time"
	"transactionx/internal/database"
	"transactionx/internal/resources"
)

type Services interface {
	RegisterTransaction(t resources.Transaction) (resources.Transaction, error)
}

type service struct {
	db database.Client
}

func NewService(db database.Client) Services {
	return &service{
		db: db,
	}
}

func (s *service) RegisterTransaction(t resources.Transaction) (resources.Transaction, error) {
	if len(t.Description) > 50 {
		return resources.Transaction{}, nil
	}
	if _, err := time.Parse(time.RFC3339, t.Date); err != nil {
		return resources.Transaction{}, nil
	}
	if t.PruchaseAmount <= 0 {
		return resources.Transaction{}, nil
	}
	return s.db.RegisterTransaction(t)
}
