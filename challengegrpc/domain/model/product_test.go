package model_test

import (
	"testing"

	"github.com/lucasti79/desafiogrpc/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	t.Run("Should create a new product", func(t *testing.T) {
		productName := "Product Test"
		producDescription := "Product Description Test"
		productPrice := float32(10.00)
		product, err := model.NewProduct(productName, producDescription, productPrice)
		require.NoError(t, err)
		assert.NotNil(t, product)
	})
}
