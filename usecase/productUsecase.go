package usecase

import (
	"api/models"
	"api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRespository
}

func NewProductUsecase(rep repository.ProductRespository) ProductUsecase {
	return ProductUsecase{
		repository: rep,
	}
}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}
