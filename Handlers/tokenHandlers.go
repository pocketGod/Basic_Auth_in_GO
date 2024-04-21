package Handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pocketGod/Basic_Auth_in_GO/Models"
	"github.com/pocketGod/Basic_Auth_in_GO/Repositories"
	"github.com/pocketGod/Basic_Auth_in_GO/Services"
)

func TokenIssueHandler(w http.ResponseWriter, r *http.Request) {
	var user Models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedUser, err := Repositories.GetUserByUsername(user.Username)
	if err != nil || storedUser == nil || storedUser.Password != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := Services.GenerateToken(*storedUser)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func TokenValidateHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	isValid, err := Services.ValidateToken(token)
	if err != nil {
		http.Error(w, "Error validating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"valid": isValid})
}
