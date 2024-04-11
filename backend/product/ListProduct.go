package product

import (
	"encoding/json"
	"mobiles/connect"
	"net/http"
)

// // func ListProducts(w http.ResponseWriter, r *http.Request) {
// 	// Fetch all products from database
// 	db := connect.InitDB()
// 	rows, err := db.Query("SELECT * FROM products")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	// Iterate over rows and collect products
// 	var products []Product
// 	for rows.Next() {
// 		var product Product
// 		if err := rows.Scan(&product.ID, &product.Name, &product.Image, &product.Price, &product.Specs); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		products = append(products, product)
// 	}

// 	// Return products as JSON response
// 	json.NewEncoder(w).Encode(products)
// }

func ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db := connect.InitDB()
	rows, err := db.Query("SELECT id, productname, price, specs, imagesName FROM products")  
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Specs, &product.Image)
		if err != nil {
			http.Error(w, "Internal Server Error____", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
