package handellers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHellowHandller(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Inside Hellow Handeller")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Something Went Wrong", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "The value in the body is %v", d)
}
