package products

import (
	"github.com/incognito/store-api/internal/categories"
)

type Service struct {
	repo    *Repository
	catrepo *categories.Repository
}

func NewService(repo *Repository, catrepo *categories.Repository) *Service {
	return &Service{
		repo:    repo,
		catrepo: catrepo,
	}
}

// check products by name & category
func (s *Service) ProductCreate(product AddProductsReq) (string, error) {

	catexists, err := s.catrepo.CategoryIDExists(product.CategoryID)

	if err != nil {
		return "", err
	}

	if !catexists {
		return "", ErrCategoryNotFound

	}
	exists, err := s.repo.ProductExists(product)

	if err != nil {
		return "", err
	}

	if exists {
		return "", ErrDuplicateProduct
	}

	return s.repo.AddProduct(product)

}

func (s *Service) GetProduct() ([]Productresponse, error) {

	return s.repo.GetProducts()

}

func (s *Service) GetProductById(id string) (Productresponse, error) {

	return s.repo.GetProductById(id)

}

//update tha productss

func (s *Service) UpdateProduct(updateproduct AddProductsReq, id string) (Productresponse, error) {

	exists, err := s.repo.ProductIdExists(id)

	if err != nil {
		return Productresponse{}, err
	}

	if !exists {
		return Productresponse{}, ErrProductNotFound
	}

	catexists, err := s.catrepo.CategoryIDExists(updateproduct.CategoryID)

	if err != nil {
		return Productresponse{}, err
	}

	if !catexists {
		return Productresponse{}, ErrCategoryNotFound

	}

	duplicateExist, err := s.repo.ProductExistsForUpdate(updateproduct, id)

	if err != nil {
		return Productresponse{}, err
	}

	if duplicateExist {
		return Productresponse{}, ErrDuplicateProduct
	}

	return s.repo.UpdateProduct(updateproduct, id)

}

func (s *Service) Deleteproduct(id string) (int64, error) {

	exists, err := s.repo.ProductIdExists(id)

	if err != nil {
		return 0, err
	}

	if !exists {

		return 0, err
	}
	return s.repo.Deleteproduct(id)

}
