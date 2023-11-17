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
		fmt.Println("1")
		http.Error(rw, "Unable to fetch products", http.StatusInternalServerError)
	}
	productsJSON, errParsing := json.Marshal(lp)
	if errParsing != nil {
		fmt.Println("2")
		http.Error(rw, "Failed to serialize products", http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")

	rw.Write(productsJSON)
}
