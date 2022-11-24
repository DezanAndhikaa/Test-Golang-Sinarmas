package usecase

import "test-golang/domain"

func (i impl) GetByID(id int) (domain.Product, error) {
	product, err := i.repo.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}

	if product == nil {
		return domain.Product{}, nil
	}

	return *product, err
}
