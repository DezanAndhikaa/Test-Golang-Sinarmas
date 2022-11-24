package product

import "test-golang/domain"

//go:generate mockgen -destination=./mocks/repository/mock_repository.go -package=mock_repository -source=repository.go

type Repository interface {
	Get(page, limit, size int) ([]domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Create(data domain.Product) (*domain.Product, error)
	Update(id int, newProduct domain.Product) error
	Delete(id int) error
}
