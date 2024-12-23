package product

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/server"
)

type Handler struct {
	productService *Service
	respService    *server.ResponseService
}

func NewHandler(productService *Service, respService *server.ResponseService) *Handler {
	return &Handler{
		productService: productService,
		respService:    respService,
	}
}

func (h *Handler) InitializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /products", h.CreateProduct)
	mux.HandleFunc("GET /products/{id}", h.GetProductByID)
	mux.HandleFunc("PATCH /products/{id}", h.UpdateProductByID)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	msgErr := "failed to create product: %w"
	dto, err := NewEntryDto(r.Body)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, wrappedErr.Error(), http.StatusBadRequest, wrappedErr)
		return
	}

	product, err := h.productService.Create(dto)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		if errors.Is(err, db.ErrDBDuplicateKey) {
			h.respService.CreateResponseError(w, "Already exist product", http.StatusConflict, wrappedErr)
			return
		}

		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
		return
	}

	response, err := NewResponseCreateDto(product)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
		return
	}

	h.respService.CreateResponseCreated(w, response)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	msgErr := "failed to create product: %w"
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, wrappedErr)
		return
	}

	product, err := h.productService.GetByID(id)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		if errors.Is(err, db.ErrDBNotFound) {
			h.respService.CreateResponseError(w, "Not found", http.StatusNotFound, wrappedErr)
			return
		}
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
		return
	}

	response, err := NewResponseReadDto(product)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
		return
	}

	h.respService.CreateResponseOK(w, response)
}

func (h *Handler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	msgErr := "failed to create update: %w"
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, wrappedErr)
		return
	}

	dto, err := NewEntryDto(r.Body)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, wrappedErr.Error(), http.StatusBadRequest, wrappedErr)
		return
	}

	_, err = h.productService.Update(id, dto)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		if errors.Is(err, db.ErrDBDuplicateKey) {
			h.respService.CreateResponseError(w, "Already exist product with same name", http.StatusConflict, wrappedErr)
			return
		}
		if errors.Is(err, db.ErrDBNotFound) {
			h.respService.CreateResponseError(w, "Not found", http.StatusNotFound, wrappedErr)
			return
		}
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
	}

	h.respService.CreateResponseNoContent(w)
}
