package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/lucasti79/desafiogrpc/domain/model"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (r ProductRepositoryDb) RegisterProduct(product *model.Product) (*model.Product, error) {
	err := r.Db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductRepositoryDb) FindProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := r.Db.Find(&products).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return products, nil
}
