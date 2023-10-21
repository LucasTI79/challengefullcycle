package mocks

import (
	"github.com/lucasti79/desafiogrpc/domain/model"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r ProductRepositoryMock) RegisterProduct(product *model.Product) (*model.Product, error) {
	args := r.Called()
	return args.Get(0).(*model.Product), args.Error(1)
}

func (r ProductRepositoryMock) FindProducts() ([]*model.Product, error) {
	args := r.Called()
	return args.Get(0).([]*model.Product), args.Error(1)
}
