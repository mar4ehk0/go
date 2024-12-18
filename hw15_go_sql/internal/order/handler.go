package order

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/mar4ehk0/go/hw15_go_sql/pkg/server"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /orders", h.Create)
	mux.HandleFunc("GET /orders/{id}", h.GetById)
	// mux.HandleFunc("PATCH /products/{id}", h.UpdateProductById)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := NewCreateDto(r.Body)

	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
			os.Stdout.Write([]byte("Can't create order, wrong input data.\n"))
			return
		}
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	order, err := h.service.Create(dto)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error() + "\n"))
		return
	}

	data, err := json.Marshal(map[string]int{"id": order.Id})

	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		os.Stdout.Write([]byte(err.Error()))
		return
	}

	server.CreateResponse(w, []byte(data), http.StatusCreated)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {

}
