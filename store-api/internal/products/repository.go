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
		INSERT into products (name, description, price, stock_quantity, category_id, image_url)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	var id string

	err := r.db.QueryRow(query, addproduct.Name, addproduct.Description, addproduct.Price,
		addproduct.StockQuantity, addproduct.CategoryID, addproduct.ImageURL).Scan(&id)

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
		AND category_id = $2
	`
	var count int

	err := r.db.QueryRow(query, addproduct.Name, addproduct.CategoryID).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (r *Repository) GetProducts() ([]Productresponse, error) {

	query := `
		SELECT id, name, description, price, stock_quantity, category_id, image_url
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

		if err := rows.Scan(&prods.ID, &prods.Name, &prods.Description, &prods.Price, &prods.StockQuantity, &prods.CategoryID, &prods.ImageURL); err != nil {
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
		SELECT id, name, description, price, stock_quantity, category_id, image_url
		FROM products
		WHERE  id = $1
	`

	var prod Productresponse

	err := r.db.QueryRow(query, id).Scan(&prod.ID, &prod.Name, &prod.Description, &prod.Price, &prod.StockQuantity, &prod.CategoryID, &prod.ImageURL)

	switch err {

	case nil:
		return prod, nil

	case sql.ErrNoRows:

		return Productresponse{}, ErrProductNotFound

	default:
		return Productresponse{}, err
	}

}

func (r *Repository) UpdateProduct(updateproduct AddProductsReq, id string) (Productresponse, error) {

	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, stock_quantity = $4, category_id = $5, image_url = $6
		WHERE id = $7
		RETURNING id, name, description, price, stock_quantity, category_id, image_url
	`

	var updatedproduct Productresponse

	err := r.db.QueryRow(query, updateproduct.Name, updateproduct.Description, updateproduct.Price,
		updateproduct.StockQuantity, updateproduct.CategoryID, updateproduct.ImageURL, id).Scan(&updatedproduct.ID, &updatedproduct.Name, &updatedproduct.Description, &updatedproduct.Price, &updatedproduct.StockQuantity, &updatedproduct.CategoryID, &updatedproduct.ImageURL)

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
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
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
