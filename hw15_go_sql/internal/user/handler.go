package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
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
	mux.HandleFunc("GET /users/{id}", h.GetById)
	mux.HandleFunc("PATCH /users/{id}", h.UpdateById)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := NewCreateDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
			os.Stdout.Write([]byte("Can't create user, wrong input data.\n"))
			return
		}
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	user, err := h.service.Create(dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			server.CreateResponse(w, []byte("Already exist user"), http.StatusConflict)
			os.Stdout.Write([]byte("Can't create user, already exist user.\n"))
			return
		}

		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error() + "\n"))
		return
	}

	data, err := json.Marshal(map[string]int{"id": user.Id})
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error()))
		return
	}

	server.CreateResponse(w, []byte(data), http.StatusCreated)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetById(id)
	if err != nil {
		if errors.Is(err, db.ErrDBNotFound) {
			server.CreateResponse(w, []byte("Not found"), http.StatusNotFound)
			os.Stdout.Write([]byte(err.Error() + "\n"))
			return
		}
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error() + "\n"))
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error()))
		return
	}

	server.CreateResponse(w, []byte(data), http.StatusCreated)
}

func (h *Handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		return
	}

	dto, err := NewUpdateDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
			os.Stdout.Write([]byte("Can't update user, wrong input data.\n"))
			return
		}
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateById(id, dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			server.CreateResponse(w, []byte("Already exist user with same email"), http.StatusConflict)
			os.Stdout.Write([]byte("Can't update user, Already exist user with same email.\n"))
			return
		}

		if errors.Is(err, db.ErrDBNotFound) {
			server.CreateResponse(w, []byte("Not found"), http.StatusNotFound)
			os.Stdout.Write([]byte(err.Error() + "\n"))
			return
		}

		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error() + "\n"))
		return
	}

	server.CreateResponse(w, []byte(""), http.StatusNoContent)
}
