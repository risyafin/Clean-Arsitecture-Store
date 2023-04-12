package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/products", jwtMiddleware(createProduct)).Methods("POST")
	r.HandleFunc("/products", jwtMiddleware(getAllProduct)).Methods("GET")
	r.HandleFunc("/products/{id}", jwtMiddleware(getProduct)).Methods("GET")
	r.HandleFunc("/products/{id}", jwtMiddleware(updateProduct)).Methods("PUT")
	r.HandleFunc("/products/{id}", jwtMiddleware(deleteProduct)).Methods("DELETE")

	r.HandleFunc("/categories", jwtMiddleware(createCategorie)).Methods("POST")
	r.HandleFunc("/categories", jwtMiddleware(getAllCategorie)).Methods("GET")
	r.HandleFunc("/categories/{id}", jwtMiddleware(getCategorie)).Methods("GET")
	r.HandleFunc("/categories/{id}", jwtMiddleware(updateCategorie)).Methods("PUT")
	r.HandleFunc("/categories/{id}", jwtMiddleware(deleteCategorie)).Methods("DELETE")

	fmt.Println("starting web server at localhost: 8080")
	http.ListenAndServe(":8080", r)
}
