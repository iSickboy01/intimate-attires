package products

import "time"

type Products struct {
	ID            string
	Name          string
	Description   string
	Price         float64
	StockQuantity int
	Category      string
	ImageURL      string
	InsertedAt    time.Time
	UpdatedAt     time.Time
}

type AddProductsReq struct {
	Name          string
	Description   string
	Price         float64
	StockQuantity int
	Category      string
	ImageURL      string
}

type Productresponse struct {
	ID            string
	Name          string
	Description   string
	Price         float64
	StockQuantity int
	Category      string
	ImageURL      string
}
