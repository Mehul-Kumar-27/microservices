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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.fetchProducts(rw, r)

		return
	}

	if r.Method == http.MethodPost {
		p.l.Println("In the post request")
		p.addNewProduct(rw, r)
		return
	}

	// if r.Method == http.MethodPut {
	// }

}

func (p *Products) fetchProducts(rw http.ResponseWriter, h *http.Request) {
	lp, err := data.GetProducts()

	if err != nil {
		fmt.Println("1")
		http.Error(rw, "Unable to fetch products", http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")

	e := json.NewEncoder(rw)

	error := e.Encode(lp)

	if error != nil {
		fmt.Println("1")
		http.Error(rw, "Unable to convert the products to json data", http.StatusInternalServerError)
	}
}

func (p *Products) addNewProduct(rw http.ResponseWriter, r *http.Request) {
	response, err := data.AddProducts(r.Body)

	p.l.Println(response)

	if err != nil {
		fmt.Println("2")
		http.Error(rw, "Error Occured", http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")

	e := json.NewEncoder(rw)

	error := e.Encode(response)

	if error != nil {
		fmt.Println("1")
		http.Error(rw, "Unable to convert the products to json data", http.StatusInternalServerError)
	}

}
