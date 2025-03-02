package repository

import (
	"api/models"
	"database/sql"
)

const (
	getAll = "SELECT id,name,price,quantity,description FROM products;"
)

type ProductRespository struct {
	connection *sql.DB
}

func NewProductRespository(connection *sql.DB) ProductRespository {
	return ProductRespository{connection: connection}
}

func (pr *ProductRespository) GetProducts() ([]models.Product, error) {
	query := getAll
	rows, err := pr.connection.Query(query)
	if err != nil {
		return []models.Product{}, err
	}
	var products []models.Product

	for rows.Next() {
		var product models.Product // Criar um novo objeto em cada iteração
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Description,
		)

		if err != nil {
			return []models.Product{}, err
		}
		products = append(products, product)
	}
	rows.Close()
	return products, nil
}
