package product

import (
	"errors"
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
	dto, err := NewEntryDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValid) {
			h.respService.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest, err)
			return
		}
		h.respService.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest, err)
		return
	}

	product, err := h.productService.Create(dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			h.respService.CreateResponse(w, []byte("Already exist product"), http.StatusConflict, err)
			return
		}

		h.respService.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError, err)
		return
	}

	data, err := NewResponseCreateDto(product)
	if err != nil {
		h.respService.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError, err)
		return
	}

	h.respService.CreateResponse(w, data, http.StatusCreated, nil)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		h.respService.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest, err)
		return
	}

	product, err := h.productService.GetByID(id)
	if err != nil {
		if errors.Is(err, db.ErrDBNotFound) {
			h.respService.CreateResponse(w, []byte("Not found"), http.StatusNotFound, err)
			return
		}
		h.respService.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError, err)
		return
	}

	data, err := NewResponseReadDto(product)
	if err != nil {
		h.respService.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError, err)
		return
	}

	h.respService.CreateResponse(w, data, http.StatusCreated, nil)
}

func (h *Handler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		h.respService.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest, err)
		return
	}

	dto, err := NewEntryDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValid) {
			h.respService.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest, err)
			return
		}
		h.respService.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError, err)
		return
	}

	_, err = h.productService.Update(id, dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			h.respService.CreateResponse(w, []byte("Already exist product with same name"), http.StatusConflict, err)
			return
		}
		if errors.Is(err, db.ErrDBNotFound) {
			h.respService.CreateResponse(w, []byte("Not found"), http.StatusNotFound, err)
			return
		}
		h.respService.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError, err)
	}

	h.respService.CreateResponse(w, []byte{}, http.StatusNoContent, nil)
}
