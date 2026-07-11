package products

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

// repository constructor
func NewRepository(db *sql.DB) *Repository {

	return &Repository{db: db}
}

func (r *Repository) AddProduct(addproduct AddProductsReq) (string, error) {

	query := `
		INSERT into products (name, description, price, stock_quantity, category, image_url)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	var id string

	err := r.db.QueryRow(query, addproduct.Name, addproduct.Description, addproduct.Price,
		addproduct.StockQuantity, addproduct.Category, addproduct.ImageURL).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil

}

// returns true if product exists
func (r *Repository) ProductExists(addproduct AddProductsReq) (bool, error) {

	query := `
		SELECT COUNT(*) 
		FROM products 
		WHERE  name = $1
		AND category = $2
	`
	var count int

	err := r.db.QueryRow(query, addproduct.Name, addproduct.Category).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (r *Repository) GetProducts() ([]Productresponse, error) {

	query := `
		SELECT id, name, description, price, stock_quantity, category, image_url
		FROM products
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var prodlist []Productresponse

	for rows.Next() {
		var prods Productresponse

		if err := rows.Scan(&prods.ID, &prods.Name, &prods.Description, &prods.Price, &prods.StockQuantity, &prods.Category, &prods.ImageURL); err != nil {
			return nil, err
		}

		prodlist = append(prodlist, prods)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prodlist, nil
}

func (r *Repository) GetProductById(id string) (Productresponse, error) {

	query := `
		SELECT id, name, description, price, stock_quantity, category, image_url
		FROM products
		WHERE  id = $1
	`

	var prod Productresponse

	err := r.db.QueryRow(query, id).Scan(&prod.ID, &prod.Name, &prod.Description, &prod.Price, &prod.StockQuantity, &prod.Category, &prod.ImageURL)

	if err != nil {
		return Productresponse{}, err
	}

	return prod, nil
}
