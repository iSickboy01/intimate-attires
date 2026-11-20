package products

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

func (h *Handler) AddProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var addreq AddProductsReq

	err := json.NewDecoder(r.Body).Decode(&addreq)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	log.Println("za error:", addreq)

	id, err := h.serv.ProductCreate(addreq)

	if err != nil {

		if errors.Is(err, ErrDuplicateProduct) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else if errors.Is(err, ErrCategoryNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"id": id,
	})
}

func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	list, err := h.serv.GetProduct()

	if err != nil {
		http.Error(w, "failed to get products", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)

}

func (h *Handler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id") // Extract path parameter
	log.Println("product id:", id)

	list, err := h.serv.GetProductById(id)

	switch err {

	case nil:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(list)

	case ErrProductNotFound:
		http.Error(w, "Product does not exist!", http.StatusNotFound)
		return

	default:
		http.Error(w, "failed to get products", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id") // Extract path parameter
	log.Println("product id:", id)

	//check product uuid before updating product
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	var updatereq AddProductsReq

	err := json.NewDecoder(r.Body).Decode(&updatereq)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	//check category uuid before updating product
	if _, err := uuid.Parse(updatereq.CategoryID); err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	updatedprod, err := h.serv.UpdateProduct(updatereq, id)

	if err != nil {

		if errors.Is(err, ErrDuplicateProduct) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else if errors.Is(err, ErrCategoryNotFound) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if errors.Is(err, ErrProductNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedprod)

}

func (h *Handler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id") //extract the path parameter

	deletedproducts, err := h.serv.Deleteproduct(id)

	if err != nil {
		http.Error(w, "failed to delete product", http.StatusInternalServerError)
		return
	}

	if deletedproducts > 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "product doesn't exist!", http.StatusNotFound)
	}
}
