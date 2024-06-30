package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"idotno.fr/echo/services"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.ListUsers()

	if err != nil {
		http.Error(w, "Could not query users", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		log.Println("Error encoding users to JSON:", err)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	user, err := services.GetUser(uint(id))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Not Found", http.StatusNotFound)
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}
