package usecase

import (
	"api/models"
	"api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(rep repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: rep,
	}
}

func (pu *ProductUseCase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product models.Product) (models.Product, error) {
	pI, err := pu.repository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}
	product.ID = pI
	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (models.Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (pu *ProductUseCase) DeleteProduct(id int) error {
	return pu.repository.DeleteProduct(id)
}
