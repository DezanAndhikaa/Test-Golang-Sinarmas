package usecase

import (
	"test-golang/application/product"
	"test-golang/application/product/dto"
	"test-golang/domain"
)

type impl struct {
	repo product.Repository
}

func (i impl) Update(request dto.UpdateProductRequest) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (i impl) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func New(repo product.Repository) product.Usecase {
	return impl{
		repo: repo,
	}
}
