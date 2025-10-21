package product

import (
	"ecommerce/domain"
	prdctHandlr "ecommerce/rest/handlers/product"
)

type Service interface{
  prdctHandlr.Service
}

type ProductRepo interface{
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(productID int) error
	Update(p domain.Product) (*domain.Product, error)
}