package categories

import "errors"

var (
	ErrEmptyCategoryName     = errors.New("Category name must not be empty!")
	ErrDuplicateCategory     = errors.New("Category already exists!")
	ErrEmptyCategoryNotFound = errors.New("Category does not exist")
)
