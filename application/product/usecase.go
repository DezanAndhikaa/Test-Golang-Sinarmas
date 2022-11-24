package product

import (
	"test-golang/application/product/dto"
	"test-golang/domain"
)

type Usecase interface {
	Create(request dto.CreateProductRequest) (domain.Product, error)
	Update(request dto.UpdateProductRequest) (domain.Product, error)
	Delete(id int) error
	GetByID(id int) (domain.Product, error)
	Get(request dto.PaginationRequest) (dto.GetAllResponse, error)
}
