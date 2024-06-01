package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fileServer)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err!= nil {
		panic(err)
	}
}