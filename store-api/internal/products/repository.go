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

func (r *Repository) UpdateProduct(updateproduct AddProductsReq, id string) (Productresponse, error) {

	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, stock_quantity = $4, category = $5, image_url = $6
		WHERE id = $7
		RETURNING id, name, description, price, stock_quantity, category, image_url
	`

	var updatedproduct Productresponse

	err := r.db.QueryRow(query, updateproduct.Name, updateproduct.Description, updateproduct.Price,
		updateproduct.StockQuantity, updateproduct.Category, updateproduct.ImageURL, id).Scan(&updatedproduct.ID, &updatedproduct.Name, &updatedproduct.Description, &updatedproduct.Price, &updatedproduct.StockQuantity, &updatedproduct.Category, &updatedproduct.ImageURL)

	if err != nil {
		return Productresponse{}, err
	}

	return updatedproduct, nil

}

func (r *Repository) Deleteproduct(id string) (int64, error) {

	query := `
	DELETE FROM products 
	WHERE id = $1
	`

	result, err := r.db.Exec(query, id)

	if err != nil {
		return 0, nil
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return rows, nil

}

// returns true if product exists
func (r *Repository) ProductIdExists(id string) (bool, error) {

	query := `
		SELECT COUNT(*) 
		FROM products 
		WHERE  id = $1
	`
	var count int

	err := r.db.QueryRow(query, id).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil

}
