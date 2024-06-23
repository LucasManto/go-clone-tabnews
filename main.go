package main

import (
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

	http.Handle("/api/v1/status", http.HandlerFunc(status))

	fmt.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
