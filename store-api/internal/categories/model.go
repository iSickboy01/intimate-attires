package categories

import "time"

type Categories struct {
	ID         string
	Name       string
	InsertedAt time.Time
	UpdatedAt  time.Time
}

type CreateCategoriesReq struct {
	Name string
}

type CategoriesResponse struct {
	ID   string
	Name string
}
