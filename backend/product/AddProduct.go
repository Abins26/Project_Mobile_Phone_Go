package product

import (
	"fmt"
	"io"
	"mobiles/connect"
	"net/http"
	"os"
	"path/filepath"
)

type Product1 struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"imageName"`
	Price float64 `json:"price"`
	Specs string  `json:"specs"`
}

// Add a new product
func AddProduct(w http.ResponseWriter, r *http.Request) {

	// Get the file part from the request
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to receive image", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// Parse other form fields to get product data
	productname := r.FormValue("name")
	price := r.FormValue("price")
	specs := r.FormValue("specs")

	//Extract the filename of the uploaded image from the multipart form
	filename := r.MultipartForm.File["image"][0].Filename
	fmt.Println(filename)

	fileDir := "../uploads/"
	if _, err := os.Stat(fileDir); os.IsNotExist(err) { //create dir if not exist
		fmt.Println("error", err)
		os.Mkdir(fileDir, os.ModePerm)
	}

	filePath := filepath.Join(fileDir, filename) // Construct the full path for saving the uploaded file

	outFile, err := os.Create(filePath) //Create a new file
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		// If there is an error copying the file contents, return an HTTP 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error copying file contents:", err)
		return
	}
	// Insert new product into database with image name
	db := connect.InitDB()
	_, err = db.Exec("INSERT INTO products (productname, price, specs, imagesName) VALUES (?, ?, ?, ?)", productname, price, specs, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Write([]byte("successfully added"))
}
