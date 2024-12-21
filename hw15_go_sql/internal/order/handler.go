package order

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/server"
)

type Handler struct {
	orderService *Service
	respService  *server.ResponseService
}

func NewHandler(orderService *Service, respService *server.ResponseService) *Handler {
	return &Handler{orderService: orderService, respService: respService}
}

func (h *Handler) InitializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /orders", h.Create)
	mux.HandleFunc("GET /orders/{id}", h.GetByID)
	mux.HandleFunc("PUT /orders/{id}", h.Update)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	entryDto, err := NewEntryCreateDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			h.respService.CreateResponseError(w, "Not valid values", http.StatusBadRequest, err)
			return
		}
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, err)
		return
	}

	order, err := h.orderService.Create(entryDto)
	if err != nil {
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	response, err := NewResponseCreateDto(order)
	if err != nil {
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	h.respService.CreateResponseCreated(w, response)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, err)
		return
	}

	outputDto, err := h.orderService.GetByID(id)
	if err != nil {
		if errors.Is(err, db.ErrDBNotFound) {
			h.respService.CreateResponseError(w, "Not found", http.StatusNotFound, err)
			return
		}
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	response, err := NewResponseReadDto(outputDto)
	if err != nil {
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	h.respService.CreateResponseOK(w, response)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	orderID, err := strconv.Atoi(idRaw)
	if err != nil {
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, err)
		return
	}

	entryDto, err := NewEntryUpdateDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			h.respService.CreateResponseError(w, "Not valid values", http.StatusBadRequest, err)
			return
		}
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, err)
		return
	}

	_, err = h.orderService.Update(orderID, entryDto)
	if err != nil {
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	h.respService.CreateResponseNoContent(w)
}
