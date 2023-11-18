package main

import (
	"fmt"
	// "kasir/modules/categories"
	"kasir/modules/logins"
	"kasir/modules/products"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var (
// 	httpRouter Router = newMuxRouter()
// )

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// categoryRepo := categories.Repository{DB: db}
	// categoryUsecase := categories.Usecase{Repo: categoryRepo}
	// categoryHandler := categories.Handler{Usecase: categoryUsecase}

	productRepo := products.Repository{DB: db}
	productUsecase := products.Usecase{Repo: productRepo}
	productHandler := products.Handler{Usecase: productUsecase}

	loginRepo := logins.Repository{DB: db}
	loginUsecase := logins.Usecase{Repo: loginRepo}
	loginHandler := logins.Handler{Usecase: loginUsecase}

	const port string = ":8080"
	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler.Login).Methods("POST")
	r.HandleFunc("/products", jwtMiddleware(productHandler.GetAllProducts)).Methods("GET")
	r.HandleFunc("/products/{id}", jwtMiddleware(productHandler.GetProduct)).Methods("GET")
	r.HandleFunc("/products", jwtMiddleware(productHandler.CreateProduct)).Methods("POST")
	r.HandleFunc("/products/{id}", jwtMiddleware(productHandler.UpdateProduct)).Methods("PUT")
	r.HandleFunc("/products/{id}", jwtMiddleware(productHandler.DeleteProduct)).Methods("DELETE")

	// r.HandleFunc("/categories/{id}", jwtMiddleware(categoryHandler.GetCategory)).Methods("GET")
	// r.HandleFunc("/categories", jwtMiddleware(categoryHandler.GetAllCategories)).Methods("GET")
	// r.HandleFunc("/categories", jwtMiddleware(categoryHandler.CreateCategory)).Methods("POST")
	// r.HandleFunc("/categories/{id}", jwtMiddleware(categoryHandler.UpdateCategory)).Methods("PUT")
	// r.HandleFunc("/categories/{id}", jwtMiddleware(categoryHandler.DeleteCategory)).Methods("DELETE")

	fmt.Println("lohalhost:8080")
	http.ListenAndServe(port, r)
}
