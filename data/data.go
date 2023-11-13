package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Product struct {
	ID                 int
	Title              string
	Description        string
	Price              float64
	DiscountPercentage float64
	Rating             float64
	Stock              int
	Brand              string
	Category           string
	Thumbnail          string
	Images             []string
}

func GetProducts() ([]Product, error) {
	url := "https://dummyjson.com/products"

	response, err := http.Get(url)

	if err != nil {
		return nil, err

	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprint("Unable to fetch produts. Status code %d", response.StatusCode))
	}

	var products []Product

	fmt.Println(response.Body)

	json.NewDecoder(response.Body).Decode(&products)

	fmt.Println(products)

	return products, nil

}
