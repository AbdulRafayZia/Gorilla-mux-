package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request utils.Credentials
	json.NewDecoder(r.Body).Decode(&request)
	role, err := service.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthozied username", http.StatusUnauthorized)
		return

	}

	validUser, err := validation.CheckUserValidity(w, r, request)
	if !validUser {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(err)
		return

	}

	tokenString, err := service.CreateToken(request.Username, role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "error in generating toke ", http.StatusInternalServerError)
		return
	}

	token := utils.Token{
		Token: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}
