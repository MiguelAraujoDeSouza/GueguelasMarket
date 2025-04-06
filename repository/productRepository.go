package repository

import (
	"api/models"
	"database/sql"
)

const (
	getAll         = "SELECT id,name,price,quantity,description FROM products;"
	createProduct  = "INSERT INTO products(name,price,quantity,description) VALUES ($1,$2,$3,$4) RETURNING id;"
	getProductById = "SELECT id,name,price,quantity,description FROM products WHERE id = $1;"
	deleteProduct  = "DELETE FROM products WHERE id = $1;"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
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

func (pr *ProductRepository) CreateProduct(product models.Product) (int, error) {
	var id int
	query := createProduct
	q, err := pr.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	err = q.QueryRow(product.Name, product.Price, product.Quantity, product.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	q.Close()
	return id, nil
}
func (pr *ProductRepository) DeleteProduct(id int) error {
	query := deleteProduct
	q, err := pr.connection.Prepare(query)
	if err != nil {
		return err
	}
	_, err = q.Exec(id)
	if err != nil {
		return err
	}
	q.Close()
	return nil
}
func (pr *ProductRepository) GetProductById(id int) (models.Product, error) {
	query := getProductById
	q, err := pr.connection.Prepare(query)
	if err != nil {
		return models.Product{}, err
	}
	var product models.Product
	err = q.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.Description,
	)
	if err != nil {
		return models.Product{}, err
	}
	q.Close()
	return product, nil
}
