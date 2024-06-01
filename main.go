package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fileServer)

	http.Handle("/api/v1/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"key": "são acima da média"})
	}))

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err!= nil {
		panic(err)
	}
}