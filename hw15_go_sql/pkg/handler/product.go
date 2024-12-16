package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/mar4ehk0/go/hw15_go_sql/internal/product"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/server"
)

type Handler struct {
	service *product.Service
}

func NewHandler(service *product.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	///
	var dto product.Dto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	if len(dto.Name) < 1 || dto.Price < 1 {
		server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
		os.Stdout.Write([]byte("Can't create product, wrong input data.\n"))
		return
	}
	/// перепишим на валидацию через product.NewDto

	pr, err := h.service.Create(dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			server.CreateResponse(w, []byte("Already exist product"), http.StatusConflict)
			os.Stdout.Write([]byte("Can't create product, already exist product.\n"))
			return
		}

		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error() + "\n"))
		return
	}

	data, err := json.Marshal(map[string]int{"id": pr.Id})

	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error()))
		return
	}

	server.CreateResponse(w, []byte(data), http.StatusCreated)
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		return
	}

	product, err := h.service.GetById(id)
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

	data, err := json.Marshal(product)

	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error()))
		return
	}

	server.CreateResponse(w, []byte(data), http.StatusCreated)
}

func (h *Handler) UpdateProductById(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		return
	}

	dto, err := product.NewDto(r.Body)
	if err != nil {
		if errors.Is(err, product.ErrNotValid) {
			server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
			os.Stdout.Write([]byte("Can't create product, wrong input data.\n"))
			return
		}
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	_, err = h.service.Update(id, *dto)
	if err != nil {
		if errors.Is(err, db.ErrDBDuplicateKey) {
			server.CreateResponse(w, []byte("Already exist product with same name"), http.StatusConflict)
			os.Stdout.Write([]byte("Can't update product.\n"))
			return
		}

		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error() + "\n"))
	}

	server.CreateResponse(w, []byte(""), http.StatusNoContent)
}
