package handellers

import (
	"encoding/json"
	"fmt"
	"log"
	"microservice/data"
	"net/http"
	"regexp"
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

	if r.Method == http.MethodPut {
		p.l.Println("In the put section")
		p.updateAProduct(rw, r)
		return
	}

	if r.Method == http.MethodDelete {
		p.l.Println("In the delete section")
		p.deleteAProduct(rw, r)
		return

	}

	rw.WriteHeader(404)

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

func (p *Products) updateAProduct(rw http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile(`(\d+)`)
	p.l.Printf(r.URL.Path)

	subStrings := regex.FindAll([]byte(r.URL.Path), -1)

	for i, matches := range subStrings {
		p.l.Printf("The component at %v is %v", i, string(matches))
	}

	// Assuming data.UpdateProduct takes a reader and a string ID
	response, err := data.UpdateProduct(r.Body, string(subStrings[0]))

	if err != nil {
		p.l.Printf("Error updating product: %v", err)
		http.Error(rw, "Error Occurred", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Update successful. Response: %v", response)

	rw.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(rw)

	if err := e.Encode(response); err != nil {
		p.l.Printf("Error encoding response: %v", err)
		http.Error(rw, "Unable to convert the products to JSON data", http.StatusInternalServerError)
	}
}

func (p *Products) deleteAProduct(rw http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile(`(\d+)`)
	p.l.Printf(r.URL.Path)

	subStrings := regex.FindAll([]byte(r.URL.Path), -1)

	for i, matches := range subStrings {
		p.l.Printf("The component at %v is %v", i, string(matches))
	}

	// Assuming data.UpdateProduct takes a reader and a string ID
	response, err := data.DeleteProduct(r.Body, string(subStrings[0]))

	if err != nil {
		p.l.Printf("Error updating product: %v", err)
		http.Error(rw, "Error Occurred", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Delete successful. Response: %v", response)

	rw.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(rw)

	if err := e.Encode(response); err != nil {
		p.l.Printf("Error encoding response: %v", err)
		http.Error(rw, "Unable to convert the products to JSON data", http.StatusInternalServerError)
	}
}
