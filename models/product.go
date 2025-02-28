package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"nome"`
	Description string  `json:"descricao"`
	Price       float64 `json:"preco"`
	Quantity    int     `json:"quantidade"`
}
