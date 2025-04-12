package repository

import (
	"api/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db: db}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := pr.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *ProductRepository) CreateProduct(product models.Product) (int, error) {
	if err := pr.db.Create(&product).Error; err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	if err := pr.db.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepository) GetProductById(id int) (models.Product, error) {
	var product models.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}
