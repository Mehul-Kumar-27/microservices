package data

import (
	"bytes"
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

func AddProducts(r io.Reader) (Product, error) {
	url := "https://dummyjson.com/products/add"
	e := json.NewDecoder(r)

	newProduct := &Product{}
	errorProduct := &Product{}
	/// Converting the reader body to the product object
	err := e.Decode(newProduct)
	if err != nil {
		return *errorProduct, err
	}

	/// Coverting the object to the json
	payload, err := json.Marshal(newProduct)
	if err != nil {
		fmt.Println("Harr gaye yrr !!!")
		return *errorProduct, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Harr gaye yrr !!")
		return *errorProduct, fmt.Errorf("HTTP request failed with status: %d", response.StatusCode)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Harr gaye yrr !")
		return *errorProduct, fmt.Errorf("HTTP request failed with status: %d", response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(newProduct)
	if err != nil {
		fmt.Println("Harr gaye yrr")
		return *errorProduct, err
	}

	return *newProduct, nil
}

func UpdateProduct(r io.Reader, id string) (Product, error) {
	// Prepare the URL
	url := fmt.Sprintf("https://dummyjson.com/products/%s", id)

	// Create a new request
	req, err := http.NewRequest("PUT", url, r)
	if err != nil {
		return Product{}, err
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Product{}, err
	}
	defer resp.Body.Close()

	// Decode the response
	var updatedProduct Product
	err = json.NewDecoder(resp.Body).Decode(&updatedProduct)
	if err != nil {
		return Product{}, err
	}

	return updatedProduct, nil
}

func DeleteProduct(r io.Reader, id string) (Product, error) {
	// Prepare the URL
	url := fmt.Sprintf("https://dummyjson.com/products/%s", id)

	// Create a new request
	req, err := http.NewRequest("DELETE", url, r)
	if err != nil {
		return Product{}, err
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Product{}, err
	}
	defer resp.Body.Close()

	// Decode the response
	var deletedProduct Product
	err = json.NewDecoder(resp.Body).Decode(&deletedProduct)
	if err != nil {
		return Product{}, err
	}

	return deletedProduct, nil
}



