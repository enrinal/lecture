package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ListProduct struct {
	Products []Product `json:"products"`
}

var ListProductData = []Product{
	Product{ID: "1", Name: "Product 1", Price: 100},
	Product{ID: "2", Name: "Product 2", Price: 200},
	Product{ID: "3", Name: "Product 3", Price: 300},
}

func handlerProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		var id = r.FormValue("id")
		for _, product := range ListProductData {
			if product.ID == id {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(product)
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func handlerListProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(ListProduct{Products: ListProductData})
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/product", handlerProduct)
	http.HandleFunc("/products", handlerListProduct)
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
