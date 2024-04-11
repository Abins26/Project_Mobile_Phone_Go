package product

import (
	"encoding/json"
	"net/http"
	// "strconv"
	"database/sql"
	"mobiles/connect"

	// "github.com/gorilla/mux"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
	Specs string  `json:"specs"`
}

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

//LoginUser     **verified**

func LoginUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return
    }
    // Set CORS headers for the main request
    w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
    w.Header().Set("Content-Type", "application/json")


    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
	db:=connect.InitDB()
    var dbUser User
    err = db.QueryRow("SELECT role FROM users WHERE username = ? AND password = ?", user.Username, user.Password).Scan(&dbUser.Role)
    switch {
    case err == sql.ErrNoRows:
        w.WriteHeader(http.StatusUnauthorized)
        return
    case err != nil:
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    dbUser.Username = user.Username
    dbUser.Password = user.Password

    jsonResponse, err := json.Marshal(dbUser)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}





// Get all products
// func GetProducts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var products []Product
// 	rows, err := db.Query("SELECT * FROM products")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var product Product
// 		if err := rows.Scan(&product.ID, &product.Name, &product.Image, &product.Price, &product.Specs); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		products = append(products, product)
// 	}
// 	if err := rows.Err(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(products)
// }





// List all products

// func ListProducts(w http.ResponseWriter, r *http.Request) {
// 	// Fetch all products from database
// 	db:=connect.InitDB()
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


// Add a new product
// func AddProduct(w http.ResponseWriter, r *http.Request) {
// 	// Parse request body to get product data
// 	var product Product
// 	err := json.NewDecoder(r.Body).Decode(&product)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Insert new product into database
// 	db:=connect.InitDB()

// 	result, err := db.Exec("INSERT INTO products (name, image, price, specs) VALUES (?, ?, ?, ?)",
// 		product.Name, product.Image, product.Price, product.Specs)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Return success response
// 	id, _ := result.LastInsertId()
// 	json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "message": "Product added successfully"})
// }


// Edit an existing product
// func EditProduct(w http.ResponseWriter, r *http.Request) {
// 	// Parse product ID from URL path
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid product ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Parse request body to get updated product data
// 	var product Product
// 	err = json.NewDecoder(r.Body).Decode(&product)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Update product in database
// 	db:=connect.InitDB()
// 	_, err = db.Exec("UPDATE products SET name=?, image=?, price=?, specs=? WHERE id=?",
// 		product.Name, product.Image, product.Price, product.Specs, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Return success response
// 	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product updated successfully"})
// }


// Delete a product

// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	// Parse product ID from URL path
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid product ID", http.StatusBadRequest)
// 		return
// 	}
//
 	// Delete product from database
// 	db:=connect.InitDB()
// 	_, err = db.Exec("DELETE FROM products WHERE id=?", id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Return success response
// 	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product deleted successfully"})
// }


// View details of a specific product
// func ViewProduct(w http.ResponseWriter, r *http.Request) {
// 	// Parse product ID from URL path
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid product ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Fetch product details from database
// 	var product Product
// 	db:=connect.InitDB()
// 	err = db.QueryRow("SELECT * FROM products WHERE id=?", id).Scan(&product.ID, &product.Name, &product.Image, &product.Price, &product.Specs)
// 	if err != nil {
// 		http.Error(w, "Product not found", http.StatusNotFound)
// 		return
// 	}

// 	// Return product details as JSON response
// 	json.NewEncoder(w).Encode(product)
// }
