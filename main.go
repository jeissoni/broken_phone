package main

import (
	"fmt"
	"log"
	"net/http"
	"telefono_roto/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)

	http.HandleFunc("/userRegister", handler.UserRegister)

	// Start the server on port 3000
	log.Fatal(http.ListenAndServe("localhost:3000", nil))

	fmt.Println("Server is running on port 3000")
}
