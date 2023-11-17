package data

import (
	"encoding/json"
	"fmt"

	"io"
	"net/http"
)

type ProductList struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
	Skip     int       `json:"skip"`
	Limit    int       `json:"limit"`
}

type Product struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              float64  `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

func GetProducts() ([]Product, error) {
	url := "https://dummyjson.com/products"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occurred at 1")
		return nil, err
	}

	var products []Product
	var productsList ProductList

	parsingError := json.Unmarshal(body, &productsList)

	if parsingError != nil {
		fmt.Println("Error occurred at 12")
		fmt.Println(parsingError)
		return nil, err
	}

	products = productsList.Products

	return products, nil
}
