package server

import (
	"log"
	"net/http"
)

type ResponseService struct{}

func NewResponseService() *ResponseService {
	return &ResponseService{}
}

func (r ResponseService) CreateResponseOK(w http.ResponseWriter, data []byte) {
	r.setContentType(w)
	r.createResponse(w, data, http.StatusOK)
}

func (r ResponseService) CreateResponseNoContent(w http.ResponseWriter) {
	r.setContentType(w)
	w.WriteHeader(http.StatusNoContent)
}

func (r ResponseService) CreateResponseCreated(w http.ResponseWriter, data []byte) {
	r.setContentType(w)
	r.createResponse(w, data, http.StatusCreated)
}

func (r ResponseService) CreateResponseError(w http.ResponseWriter, humanError string, httpCode int, err error) {
	if err != nil {
		log.Println(err)
	}

	r.setContentType(w)
	r.createResponse(w, []byte(humanError), httpCode)
}

func (r ResponseService) setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (r ResponseService) createResponse(w http.ResponseWriter, data []byte, httpCode int) {
	w.WriteHeader(httpCode)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}
