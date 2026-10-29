package categories

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	serv *Service
}

func NewHandler(serv *Service) *Handler {
	return &Handler{serv: serv}
}

func (h *Handler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var createcatreq CreateCategoriesReq
	//var createcatresp CategoriesResponse

	err := json.NewDecoder(r.Body).Decode(&createcatreq)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createcatresp, err := h.serv.CreateCategory(createcatreq)

	//log.Println("za error:", createcatresp)

	if err != nil {
		log.Printf("failed to create category: %v", err)
		if errors.Is(err, ErrEmptyCategoryName) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if errors.Is(err, ErrDuplicateCategory) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createcatresp)

}

func (h *Handler) GetCategoryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cats, err := h.serv.GetCategory()

	if err != nil {
		http.Error(w, "failed to get categories", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cats)

}

func (h *Handler) GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	cat, err := h.serv.GetCategoryById(id)

	switch err {

	case nil:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cat)
		return

	case ErrEmptyCategoryNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	var updatecat CreateCategoriesReq

	err := json.NewDecoder(r.Body).Decode(&updatecat)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedcat, err := h.serv.UpdateCategory(updatecat, id)

	switch err {

	case nil:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedcat)
		return

	case ErrEmptyCategoryNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	case ErrEmptyCategoryName:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	case ErrDuplicateCategory:
		http.Error(w, err.Error(), http.StatusConflict)
		return

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

}
