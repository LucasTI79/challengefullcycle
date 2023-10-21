package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
)

var (
	ProductNotFound = errors.New("product not found")
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductRepositoryInterface interface {
	RegisterProduct(product *Product) (*Product, error)
	FindProducts() ([]*Product, error)
}

type Product struct {
	Base        `valid:"required"`
	Name        string  `gorm:"column:name;type:varchar(50)" json:"name,omitempty" valid:"notnull"`
	Description string  `gorm:"column:description;type:varchar(255)" json:"description,omitempty" valid:"notnull"`
	Price       float32 `gorm:"column:price" json:"price,omitempty" valid:"notnull"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product.CreatedAt = time.Now()

	err := product.isValid()
	if err != nil {
		return nil, err
	}
	return &product, nil
}
