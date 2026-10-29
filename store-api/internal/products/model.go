package products

import "time"

type Products struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	StockQuantity int       `json:"stock_quantity"`
	CategoryID    string    `json:"category_id"`
	ImageURL      string    `json:"image_url"`
	InsertedAt    time.Time `json:"inserted_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AddProductsReq struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	CategoryID    string  `json:"category_id"`
	ImageURL      string  `json:"image_url"`
}

type Productresponse struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	CategoryID    string  `json:"category_id"`
	ImageURL      string  `json:"image_url"`
}
