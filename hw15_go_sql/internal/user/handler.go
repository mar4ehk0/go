package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/server"
)

type Handler struct {
	userService *Service
	respService *server.ResponseService
}

func NewHandler(userService *Service, respService *server.ResponseService) *Handler {
	return &Handler{userService: userService, respService: respService}
}

func (h *Handler) InitializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", h.Create)
	mux.HandleFunc("GET /users/{id}", h.GetByID)
	mux.HandleFunc("PATCH /users/{id}", h.UpdateByID)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := NewEntryCreateDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			h.respService.CreateResponseError(w, "Not valid values", http.StatusBadRequest, err)
			return
		}
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.Create(dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			h.respService.CreateResponseError(w, "Already exist user", http.StatusConflict, err)
			return
		}

		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	response, err := NewResponseCreateDto(user)
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

	user, err := h.userService.GetByID(id)
	if err != nil {
		if errors.Is(err, db.ErrDBNotFound) {
			h.respService.CreateResponseError(w, "Not found", http.StatusNotFound, err)
			return
		}
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	response, err := NewResponseReadDto(user)
	if err != nil {
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	h.respService.CreateResponseOK(w, response)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, err)
		return
	}

	dto, err := NewEntryUpdateDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			h.respService.CreateResponseError(w, "Not valid values", http.StatusBadRequest, err)
			return
		}
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, err)
		return
	}

	err = h.userService.UpdateByID(id, dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			h.respService.CreateResponseError(w, "Already exist user with same email", http.StatusConflict, err)
			return
		}

		if errors.Is(err, db.ErrDBNotFound) {
			h.respService.CreateResponseError(w, "Not found", http.StatusNotFound, err)
			return
		}

		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	h.respService.CreateResponseNoContent(w)
}
