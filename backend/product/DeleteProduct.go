package product

import (
	"encoding/json"
	"mobiles/connect"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Parse product ID from URL path
	params := mux.Vars(r)
	productid, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Retrieve the image filename from the database for the product being deleted
	var imageName string
	db := connect.InitDB()
	err = db.QueryRow("SELECT imagesName FROM products WHERE id=?", productid).Scan(&imageName)
	if err != nil {
		http.Error(w, "Failed to retrieve image filename", http.StatusInternalServerError)
		return
	}

	// Delete the corresponding image file from the local directory
	imagePath := "../uploads/" + imageName // Update the path to match your setup
	err = os.Remove(imagePath)
	if err != nil {
		http.Error(w, "Failed to delete image file", http.StatusInternalServerError)
		return
	}

	// Delete product from the database
	_, err = db.Exec("DELETE FROM products WHERE id=?", productid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product deleted successfully"})
}

// Delete a product
// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	// Parse product ID from URL path
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid product ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Delete product from database
// 	db:=connect.InitDB()
// 	_, err = db.Exec("DELETE FROM products WHERE id=?", id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Return success response
// 	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product deleted successfully"})
// }
