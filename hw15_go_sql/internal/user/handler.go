package user

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/server"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
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
			server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	user, err := h.service.Create(dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			server.CreateResponse(w, []byte("Already exist user"), http.StatusConflict)
			log.Println(err.Error())
			return
		}

		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	data, err := NewResponseCreateDto(user)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	server.CreateResponse(w, data, http.StatusCreated)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, db.ErrDBNotFound) {
			server.CreateResponse(w, []byte("Not found"), http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	data, err := NewResponseReadDto(user)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	server.CreateResponse(w, data, http.StatusCreated)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	dto, err := NewEntryUpdateDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateByID(id, dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			server.CreateResponse(w, []byte("Already exist user with same email"), http.StatusConflict)
			log.Println(err.Error())
			return
		}

		if errors.Is(err, db.ErrDBNotFound) {
			server.CreateResponse(w, []byte("Not found"), http.StatusNotFound)
			log.Println(err.Error())
			return
		}

		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	server.CreateResponse(w, []byte(""), http.StatusNoContent)
}
