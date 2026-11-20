package categories

import "time"

type Categories struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	InsertedAt time.Time `json:"inserted_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateCategoriesReq struct {
	Name string `json:"name"`
}

type CategoriesResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
