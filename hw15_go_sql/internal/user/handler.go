package user

import (
	"errors"
	"fmt"
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
	msgErr := "failed to create user: %w"

	dto, err := NewEntryCreateDto(r.Body)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, wrappedErr.Error(), http.StatusBadRequest, wrappedErr)
		return
	}

	user, err := h.userService.Create(dto)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		if errors.Is(wrappedErr, db.ErrDBDuplicateKey) {
			h.respService.CreateResponseError(w, "Already exist user", http.StatusConflict, wrappedErr)
			return
		}

		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
		return
	}

	response, err := NewResponseCreateDto(user)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
		return
	}

	h.respService.CreateResponseCreated(w, response)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	msgErr := "failed to getbyid user: %w"

	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, wrappedErr)
		return
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		if errors.Is(wrappedErr, db.ErrDBNotFound) {
			h.respService.CreateResponseError(w, "Not found", http.StatusNotFound, wrappedErr)
			return
		}
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
		return
	}

	response, err := NewResponseReadDto(user)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, wrappedErr)
		return
	}

	h.respService.CreateResponseOK(w, response)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	msgErr := "failed to updatedbyid user: %w"

	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, err.Error(), http.StatusBadRequest, wrappedErr)
		return
	}

	entryDto, err := NewEntryUpdateDto(r.Body)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		h.respService.CreateResponseError(w, wrappedErr.Error(), http.StatusBadRequest, wrappedErr)
		return
	}

	err = h.userService.UpdateByID(id, entryDto)
	if err != nil {
		wrappedErr := fmt.Errorf(msgErr, err)
		if errors.Is(wrappedErr, db.ErrDBDuplicateKey) {
			h.respService.CreateResponseError(w, "Already exist user with same email", http.StatusConflict, wrappedErr)
			return
		}

		if errors.Is(wrappedErr, db.ErrDBNotFound) {
			h.respService.CreateResponseError(w, "Not found", http.StatusNotFound, wrappedErr)
			return
		}

		h.respService.CreateResponseError(w, "Something went wrong", http.StatusInternalServerError, err)
		return
	}

	h.respService.CreateResponseNoContent(w)
}
