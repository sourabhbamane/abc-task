package controllers

import (
	"encoding/json"
	"healthclub/error"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(error.Success{Message: "welcome to glofox managed studio"})
}
