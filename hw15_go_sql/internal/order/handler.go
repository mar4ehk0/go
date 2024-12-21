package order

import (
	"errors"
	"log"
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
			server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	order, err := h.orderService.Create(entryDto)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	response, err := NewResponseCreateDto(order)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	server.CreateResponse(w, response, http.StatusCreated)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		return
	}

	outputDto, err := h.orderService.GetByID(id)
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

	response, err := NewResponseReadDto(outputDto)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	server.CreateResponse(w, response, http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idRaw := r.PathValue("id")

	orderID, err := strconv.Atoi(idRaw)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		return
	}

	entryDto, err := NewEntryUpdateDto(r.Body)
	if err != nil {
		if errors.Is(err, ErrNotValidRequest) {
			server.CreateResponse(w, []byte("Not valid values"), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		server.CreateResponse(w, []byte(err.Error()), http.StatusBadRequest)
		return
	}

	_, err = h.orderService.Update(orderID, entryDto)
	if err != nil {
		server.CreateResponse(w, []byte("Something went wrong"), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	server.CreateResponse(w, []byte{}, http.StatusNoContent)
}
