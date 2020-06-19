package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// register route handler
	router.HandleFunc("/users/verify", verifyUserHandler).Methods("PUT")

	port := 8000

	fmt.Printf("staring go-api-example on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	log.Fatal(err)
}
