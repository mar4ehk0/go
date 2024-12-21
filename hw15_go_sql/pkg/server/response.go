package server

import (
	"log"
	"net/http"
)

type ResponseService struct{}

func NewResponseService() *ResponseService {
	return &ResponseService{}
}

func (r ResponseService) Response(w http.ResponseWriter, data []byte, httpCode int, err error) {
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func CreateResponse(w http.ResponseWriter, data []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}
