package categories

import (
	"strings"
)

type ProductChecker interface {
	CategoryHasProducts(categoryid string) (bool, error)
}

type Service struct {
	repo     *Repository
	prodrepo ProductChecker
}

func NewService(repo *Repository, checker ProductChecker) *Service {
	return &Service{
		repo:     repo,
		prodrepo: checker,
	}
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

	return s.repo.GetCategory()

}

func (s *Service) GetCategoryById(id string) (CategoriesResponse, error) {

	return s.repo.GetCategoryById(id)

}

func (s *Service) UpdateCategory(updatecategory CreateCategoriesReq, id string) (CategoriesResponse, error) {

	updatecategory.Name = strings.TrimSpace(updatecategory.Name)

	if updatecategory.Name == "" {
		return CategoriesResponse{}, ErrEmptyCategoryName
	}

	exist, err := s.repo.CategoryExistsForUpdate(updatecategory, id)

	if err != nil {
		return CategoriesResponse{}, err
	}

	if exist {
		return CategoriesResponse{}, ErrDuplicateCategory
	}

	return s.repo.UpdateCategory(updatecategory, id)

}

func (s *Service) Deletecategory(id string) (int64, error) {

	exists, err := s.repo.CategoryIDExists(id)

	if err != nil {
		return 0, err
	}

	if !exists {
		return 0, ErrEmptyCategoryNotFound
	}

	cathasprod, err := s.prodrepo.CategoryHasProducts(id)

	if err != nil {
		return 0, err
	}

	if cathasprod {
		return 0, ErrCategoryHasProducts
	}
	return s.repo.Deletecategory(id)

}
