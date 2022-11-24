package usecase

import (
	"test-golang/application/product/dto"
	"test-golang/domain"
)

func (i impl) Create(request dto.CreateProductRequest) (domain.Product, error) {
	newData := domain.Product{
		Name: request.Name,
		Qty:  request.Qty,
	}

	data, err := i.repo.Create(newData)
	if err != nil {
		return domain.Product{}, err
	}

	return *data, nil
}
