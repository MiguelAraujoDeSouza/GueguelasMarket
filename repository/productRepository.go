package repository

import (
	"api/models"
	"database/sql"
)

type ProductRespository struct {
	connection *sql.DB
}

func NewProductRespository(connection *sql.DB) ProductRespository {
	return ProductRespository{connection: connection}
}

func (pr *ProductRespository) GetProducts() ([]models.Product, error) {
	query := "SELECT * FROM products;"
	rows, err := pr.connection.Query(query)
	if err != nil {
		return []models.Product{}, err
	}
	var products []models.Product
	var product models.Product

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Quantity,
			&product.Description,
			&product.Price,
		)

		if err != nil {
			return []models.Product{}, err
		}
		products = append(products, product)
	}
	rows.Close()
	return products, nil
}
