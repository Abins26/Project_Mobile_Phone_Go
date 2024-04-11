package main

import (
	// "database/sql"
	"fmt"
	"log"
	"mobiles/connect"
	"mobiles/product"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Product model

// var db *sql.DB

// Initialize database connection

// Main function
func main() {
	// Initialize router
	r := mux.NewRouter()

	// Initialize database
	connect.InitDB()

	// Route Handlers / Endpoints
	// r.HandleFunc("/api/products", product.GetProducts).Methods("GET")
	//login
	r.HandleFunc("/login", product.LoginUser)
	// .Methods("GET")

	// API endpoints for product management
	r.HandleFunc("/addproducts", product.AddProduct).Methods("POST")
	// r.HandleFunc("/api/products/edit/{id}", product.EditProduct).Methods("PUT")
	r.HandleFunc("/del_products/delete/{id}", product.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/list_products", product.ListProducts).Methods("GET")
	// r.HandleFunc("/api/products/{id}", product.ViewProduct).Methods("GET")

	// Start server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
