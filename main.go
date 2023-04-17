package main

import (
	"kasir/modules/categories"
	"kasir/modules/logins"
	"kasir/modules/products"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	httpRouter Router = newMuxRouter()
)

func main() {
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	categoryRepo := categories.Repository{DB: db}
	categoryUsecase := categories.Usecase{Repo: categoryRepo}
	categoryHandler := categories.Handler{Usecase: categoryUsecase}

	productRepo := products.Repository{DB: db}
	productUsecase := products.Usecase{Repo: productRepo}
	productHandler := products.Handler{Usecase: productUsecase}

	loginRepo := logins.Repository{DB: db}
	loginUsecase := logins.Usecase{Repo: loginRepo}
	loginHandler := logins.Handler{Usecase: loginUsecase}

	const port string = ":8080"

	httpRouter.POST("/login", loginHandler.Login)
	httpRouter.GET("/products", jwtMiddleware(productHandler.GetAllProducts))
	httpRouter.GET("/products/{id}", jwtMiddleware(productHandler.GetProduct))
	httpRouter.POST("/products", jwtMiddleware(productHandler.CreateProduct))
	httpRouter.PUT("/products/{id}", jwtMiddleware(productHandler.UpdateProduct))
	httpRouter.DELETE("/products/{id}", jwtMiddleware(productHandler.DeleteProduct))

	httpRouter.GET("/categories/{id}", jwtMiddleware(categoryHandler.GetCategory))
	httpRouter.GET("/categories", jwtMiddleware(categoryHandler.GetAllCategories))
	httpRouter.POST("/categories", jwtMiddleware(categoryHandler.CreateCategory))
	httpRouter.PUT("/categories/{id}", jwtMiddleware(categoryHandler.UpdateCategory))
	httpRouter.DELETE("/categories/{id}", jwtMiddleware(categoryHandler.DeleteCategory))

	httpRouter.SERVE(port)
}
