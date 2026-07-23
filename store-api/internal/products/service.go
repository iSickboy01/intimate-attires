package products

import (
	"log"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// check products by name & category
func (s *Service) ProductCreate(product AddProductsReq) (string, error) {

	exists, err := s.repo.ProductExists(product)

	if err != nil {
		return "", err
	}

	if exists == true {
		log.Println("Product already exists!")
		return "product exists", err
	} else {
		id, err := s.repo.AddProduct(product)

		if err != nil {
			return "", err

		}
		return id, nil
	}

}

func (s *Service) GetProduct() ([]Productresponse, error) {

	prods, err := s.repo.GetProducts()

	if err != nil {
		return nil, err
	}

	return prods, nil

}

func (s *Service) GetProductById(id string) (Productresponse, error) {

	prod, err := s.repo.GetProductById(id)

	if err != nil {
		return Productresponse{}, err
	}

	return prod, nil

}

func (s *Service) UpdateProduct(updateproduct AddProductsReq, id string) (string, error) {

	id, err := s.repo.UpdateProduct(updateproduct, id)

	if err != nil {
		return "", err
	}

	return id, nil

}
