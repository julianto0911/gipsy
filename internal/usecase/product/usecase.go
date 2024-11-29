package ucproduct

import (
	"app/internal/repository"
)

func NewProductUseCase(product repository.RProduct) ProductUC {
	return &productUC{
		product: product,
	}
}

type ProductUC interface {
	Create(input InputProduct) (*repository.EProduct, error)
}

type productUC struct {
	product repository.RProduct
}

func (uc *productUC) Create(input InputProduct) (*repository.EProduct, error) {
	return uc.product.Create(input.Name)

}
