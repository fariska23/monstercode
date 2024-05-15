package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"monstercode/app"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/products.json", productsJSONHandler) // JSON endpoint
	http.HandleFunc("/api/v1/products.xml", productsXMLHandler)   // XML endpoint
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func productsJSONHandler(w http.ResponseWriter, r *http.Request) {
	products := []app.Product{
		createProduct(1, "Product 1", 50000, 2022),
		createProduct(2, "Product 2", 75000, 2023),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func productsXMLHandler(w http.ResponseWriter, r *http.Request) {
	products := []app.Product{
		createProduct(1, "Product 1", 50000, 2022),
		createProduct(2, "Product 2", 75000, 2023),
	}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(products)
}

// Helper function to create products
func createProduct(id int, name string, price int64, year int) app.Product {
	prod := app.Product{
		Price:    price,
		Category: &app.Category{},
	}
	prod.SetIDAndName(id, name)
	prod.SetYear(year)
	return prod
}
