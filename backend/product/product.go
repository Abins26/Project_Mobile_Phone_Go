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