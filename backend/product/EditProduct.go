package product

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mobiles/connect"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

func EditProduct(w http.ResponseWriter, r *http.Request) {
	var imgUrl string
	db := connect.InitDB()

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	fmt.Println("vars", err)
	if err != nil {
		http.Error(w, "Invalid mobile phone ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	specs := r.FormValue("specs")
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	fmt.Println(err)

	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	fmt.Println(err)

	if err != nil {
		http.Error(w, "Failed to get image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = db.QueryRow("select imagesName from products where id=?", id).Scan(&imgUrl)
	if err != nil {
		log.Print(err)
	}
	oldimageName := path.Base(imgUrl)

	oldimagePath := "../uploads/" + oldimageName
	err = os.Remove(oldimagePath)
	if err != nil {
		fmt.Println("failed to delete  error ", err)
		http.Error(w, "Failed to delete image file", http.StatusInternalServerError)
		return
	}

	imagename := handler.Filename
	imagePath := "../uploads/" + handler.Filename
	outputFile, err := os.Create(imagePath)
	fmt.Println(err)

	if err != nil {
		fmt.Println("error creating path", err)
		http.Error(w, "Failed to save image file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	fmt.Println(err)

	if err != nil {
		http.Error(w, "Failed to write image file", http.StatusInternalServerError)
		return
	}

	// Set ID from URL parameter

	_, err = db.Exec("UPDATE products SET productname = ?, specs = ?, price = ?,imagesName = ? WHERE id = ?",
		name, specs, price, imagename, id)
	if err != nil {
		http.Error(w, "Failed to update mobile phone", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product updated successfully"})
}
