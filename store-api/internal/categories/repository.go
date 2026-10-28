package categories

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

//Post category
func (r *Repository) CreateCategory(createcatreq CreateCategoriesReq) (CategoriesResponse, error) {

	query :=
		`INSERT into categories (name)
	VALUES ($1)
	RETURNING id, name
	`
	var response CategoriesResponse
	err := r.db.QueryRow(query, createcatreq.Name).Scan(&response.ID, &response.Name)

	if err != nil {
		return CategoriesResponse{}, err
	}

	return response, nil
}

//check if category exists
func (r *Repository) CategoryExists(createcatreq CreateCategoriesReq) (bool, error) {

	query :=
		`SELECT COUNT(*)
	from categories
	WHERE name = $1
	`
	var count int
	err := r.db.QueryRow(query, createcatreq.Name).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

//get all category

func (r *Repository) GetCategory() ([]CategoriesResponse, error) {

	query := `SELECT id, name
	FROM categories
	`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	catlist := make([]CategoriesResponse, 0)

	for rows.Next() {

		var cats CategoriesResponse
		if err := rows.Scan(&cats.ID, &cats.Name); err != nil {
			return nil, err
		}
		catlist = append(catlist, cats)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return catlist, nil

}

//Get category by id
func (r *Repository) GetCategoryById(id string) (CategoriesResponse, error) {

	query :=
		`SELECT id, name
	FROM categories
	WHERE id = $1
	`
	var cat CategoriesResponse
	err := r.db.QueryRow(query, id).Scan(&cat.ID, &cat.Name)

	switch err {

	case nil:
		return cat, nil

	case sql.ErrNoRows:
		return CategoriesResponse{}, ErrEmptyCategoryNotFound

	default:
		return CategoriesResponse{}, err
	}

}

//Update category
func (r *Repository) UpdateCategory(updatecategory CreateCategoriesReq, id string) (CategoriesResponse, error) {

	query :=
		`UPDATE categories
	SET name = $1
	WHERE id = $2
	RETURNING id, name
	`
	var cat CategoriesResponse
	err := r.db.QueryRow(query, updatecategory.Name, id).Scan(&cat.ID, &cat.Name)

	switch err {

	case nil:
		return cat, nil

	case sql.ErrNoRows:
		return CategoriesResponse{}, ErrEmptyCategoryNotFound

	default:
		return CategoriesResponse{}, err
	}

}
