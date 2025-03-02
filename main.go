package main

import (
	"api/controllers"
	"api/db"
	"api/repository"
	"api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ProductRepository := repository.NewProductRespository(dbConnection)

	ProdcutUsecase := usecase.NewProductUsecase(ProductRepository)

	ProductController := controllers.NewProdutController(ProdcutUsecase)

	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"health": "ok",
		})
	})

	api.GET("/products", ProductController.GetProducts)

	api.Run(":8080")
}
