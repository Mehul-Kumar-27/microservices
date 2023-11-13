package handellers

import (
	"encoding/json"
	"fmt"
	"log"
	"microservice/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProudcts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp, err := data.GetProducts()

	if err != nil {
		http.Error(rw, "Unable to fetch products", http.StatusInternalServerError)
	}
	productsJSON, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Failed to serialize products", http.StatusInternalServerError)
		return
	}

	fmt.Println(lp)

	rw.Header().Set("Content-Type", "application/json")

	rw.Write(productsJSON)
}
