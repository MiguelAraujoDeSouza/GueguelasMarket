package usecase

import (
	"api/models"
	"api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(rep repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: rep,
	}
}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product models.Product) (models.Product, error) {
	p_i, err := pu.repository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}
	product.ID = p_i
	return product, nil
}
