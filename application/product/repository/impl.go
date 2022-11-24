package repository

import (
	"gorm.io/gorm"
	"test-golang/application/product"
	"test-golang/domain"
)

type impl struct {
	db *gorm.DB
}

func (i impl) Get(page, limit, size int) ([]domain.Product, error) {
	var data []domain.Product

	err := i.db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (i impl) GetByID(id int) (*domain.Product, error) {
	var data domain.Product

	err := i.db.First(&data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &data, nil
}

func (i impl) Create(data domain.Product) (*domain.Product, error) {
	err := i.db.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (i impl) Update(id int, newProduct domain.Product) error {
	//TODO implement me
	panic("implement me")
}

func (i impl) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func New(db *gorm.DB) product.Repository {
	return impl{
		db: db,
	}
}
