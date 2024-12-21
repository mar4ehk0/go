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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func (r ResponseService) CreateResponseNoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (r ResponseService) CreateResponseCreated(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func (r ResponseService) CreateResponseError(w http.ResponseWriter, humanError string, httpCode int, err error) {
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, err = w.Write([]byte(humanError))
	if err != nil {
		log.Println(err)
	}
}
