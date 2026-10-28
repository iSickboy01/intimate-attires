package categories

import (
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCategory(createcatreq CreateCategoriesReq) (CategoriesResponse, error) {

	createcatreq.Name = strings.TrimSpace(createcatreq.Name)

	if createcatreq.Name == "" {

		return CategoriesResponse{}, ErrEmptyCategoryName
	}

	exists, err := s.repo.CategoryExists(createcatreq)

	if err != nil {
		return CategoriesResponse{}, err
	}

	if exists {
		return CategoriesResponse{}, ErrDuplicateCategory
	}
	response, err := s.repo.CreateCategory(createcatreq)

	if err != nil {
		return CategoriesResponse{}, err
	}
	return response, nil

}

func (s *Service) GetCategory() ([]CategoriesResponse, error) {

	cats, err := s.repo.GetCategory()

	if err != nil {
		return nil, err
	}

	return cats, nil

}

func (s *Service) GetCategoryById(id string) (CategoriesResponse, error) {

	cat, err := s.repo.GetCategoryById(id)

	switch err {

	case nil:
		return cat, nil

	case ErrEmptyCategoryNotFound:
		return CategoriesResponse{}, ErrEmptyCategoryNotFound

	default:
		return CategoriesResponse{}, err
	}

}

func (s *Service) UpdateCategory(updatecategory CreateCategoriesReq, id string) (CategoriesResponse, error) {

	updatecategory.Name = strings.TrimSpace(updatecategory.Name)

	if updatecategory.Name == "" {
		return CategoriesResponse{}, ErrEmptyCategoryName
	}

	cat, err := s.repo.UpdateCategory(updatecategory, id)

	switch err {

	case nil:
		return cat, nil

	case ErrEmptyCategoryNotFound:
		return CategoriesResponse{}, ErrEmptyCategoryNotFound

	default:
		return CategoriesResponse{}, err
	}

}
