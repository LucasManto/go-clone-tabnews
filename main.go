package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.development")
	if err != nil {
		panic(err)
	}

	log.SetFlags(log.LstdFlags | log.LUTC | log.Llongfile)

	fileServer := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fileServer)

	http.Handle("/api/v1/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response int
		err := Query("SELECT 1 + 1 as SUM;", &response)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{"key": "error"})
			return
		}
		log.Println(response)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"key": "são acima da média"})
	}))

	fmt.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
