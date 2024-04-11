package product

import (
	"encoding/json"
	"fmt"
	"io"
	"mobiles/connect"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func EditProduct(w http.ResponseWriter, r *http.Request) {
	// Parse product ID from URL path
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Parse form fields to get updated product data
	err = r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// Retrieve the existing image filename from the database for the product being updated
	var existingImageName string
	db := connect.InitDB()
	err = db.QueryRow("SELECT imagesName FROM products WHERE id=?", id).Scan(&existingImageName)
	if err != nil {
		http.Error(w, "Failed to retrieve existing image filename", http.StatusInternalServerError)
		return
	}

	// Delete the existing image file from the local directory
	existingImagePath := "../uploads/" + existingImageName // Update the path to match your setup
	err = os.Remove(existingImagePath)
	if err != nil {
		http.Error(w, "Failed to delete existing image file", http.StatusInternalServerError)
		return
	}

	// Upload the new image file to the local directory
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to receive image", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// // Generate a new filename for the uploaded image (you may use a unique identifier)
	// newImageName := generateUniqueImageName()

	
	file, _, err = r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to receive image", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	filename := r.MultipartForm.File["image"][0].Filename
	fmt.Println(filename)

	// Save the uploaded image file to the local directory
	newImagePath := "../uploads/" + filename // Update the path to match your setup
	outFile, err := os.Create(newImagePath)
	if err != nil {
		http.Error(w, "Failed to create new image file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Failed to save image file", http.StatusInternalServerError)
		return
	}

	// Update product details in the database
	_, err = db.Exec("UPDATE products SET imagesName=? WHERE id=?", filename, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product updated successfully"})
}
