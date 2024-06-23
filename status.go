package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	var response int
	err := query(r.Context(), "SELECT 1 + 1 as SUM;", &response)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"key": "error"})
		return
	}
	log.Println(response)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"key": "são acima da média"})
}
