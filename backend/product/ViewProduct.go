package product

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mobiles/connect"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ViewProduct(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")                                // Set allowed headers
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")                    // Set allowed methods
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Parse product ID from URL path

	params := mux.Vars(r)
	// fmt.Println(r)
	productid, err := strconv.Atoi(params["id"])

	// query := r.URL.Query()
	//productIdStr := r.URL.Query().Get("id")
	//productid, err := strconv.Atoi(productIdStr)
	// fmt.Println(productid)

	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Retrieve product details from the database
	db := connect.InitDB() // Assuming you have a function to initialize the database connection
	var product Product
	err = db.QueryRow("SELECT id, productname, price, specs ,imagesName FROM products WHERE id=?", productid).
		Scan(&product.ID, &product.Name, &product.Price, &product.Specs, &product.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Return product details as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
