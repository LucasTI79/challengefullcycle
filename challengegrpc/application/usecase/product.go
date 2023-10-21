package usecase

import "github.com/lucasti79/desafiogrpc/domain/model"

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) FindProducts() ([]*model.Product, error) {
	return p.ProductRepository.FindProducts()
}

func (p *ProductUseCase) RegisterProduct(name, description string, price float32) (*model.Product, error) {
	product := &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
	product, err := p.ProductRepository.RegisterProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
