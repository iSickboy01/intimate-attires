package products

import "errors"

var (
	ErrProductNotFound = errors.New("Product does not exist")
)
