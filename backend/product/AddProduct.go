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
// func AddProduct(w http.ResponseWriter, r *http.Request) {
// 	// Parse other form fields to get product data
// 	name := r.FormValue("name")
// 	price := r.FormValue("price")
// 	specs := r.FormValue("specs")

// 	// Parse uploaded file
// 	file, _, err := r.FormFile("image")
// 	if err != nil {
// 		http.Error(w, "Failed to parse image", http.StatusBadRequest)
// 		return
// 	}
// 	defer file.Close()

// 	// Save image name to database
// 	imageName := r.FormValue("imageName")
// 	if imageName == "" {
// 		http.Error(w, "Image name is required", http.StatusBadRequest)
// 		return
// 	}

// 	// Insert new product into database with image file name
// 	db:=connect.InitDB()
// 	result, err := db.Exec("INSERT INTO products (name, imageName, price, specs) VALUES (?, ?, ?, ?)",
// 		name, imageName, price, specs)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Return success response
// 	id, _ := result.LastInsertId()
// 	json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "message": "Product added successfully"})
// }

// // Add a new product
func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}
	// Set CORS headers for the main request
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "plain/text")

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

	// fmt.Println(productname)

	// Extract the filename from the file part headers
	filename := r.MultipartForm.File["image"][0].Filename
	fmt.Println(filename)

	fileDir := "/home/abinsms/Documents/New_Project_demo/Project_crud_mobile/uploads/"
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		os.Mkdir(fileDir, os.ModePerm)
	}

	filePath := filepath.Join(fileDir, filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	// id, _ := result.LastInsertId()
	// jsonbody := json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "message": "Product added successfully"})
	w.Write([]byte("successful"))
}
