package products

import "errors"

var (
	ErrProductNotFound  = errors.New("Product does not exist")
	ErrDuplicateProduct = errors.New("Product already exists!")
	ErrCategoryNotFound = errors.New("Category does not exist")
)
