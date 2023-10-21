package usecase_test

import (
	"testing"

	"github.com/lucasti79/desafiogrpc/application/usecase"
	"github.com/lucasti79/desafiogrpc/domain/model"
	"github.com/lucasti79/desafiogrpc/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var expectedProduct = model.Product{
	Name:        "PROD02",
	Description: "Yogurt",
	Price:       2.50,
}

func TestRegisterProduct(t *testing.T) {
	t.Run("Should register a new product", func(t *testing.T) {
		repository := mocks.ProductRepositoryMock{}
		sut := usecase.ProductUseCase{ProductRepository: &repository}
		repository.Mock.On("RegisterProduct").Return(&expectedProduct, nil)
		product, err := sut.RegisterProduct(expectedProduct.Name, expectedProduct.Description, expectedProduct.Price)
		require.NoError(t, err)
		assert.NotNil(t, product)
	})
}
