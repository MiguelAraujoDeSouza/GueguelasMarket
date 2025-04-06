package routes

import (
	"api/controllers"
	"api/db"
	"api/repository"
	"api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Conecta ao banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Inicializa o repository, usecase e controller
	productRepository := repository.NewProductRepository(dbConnection)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productController := controllers.NewProdutoController(productUsecase)

	// Define as rotas
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
	})

	router.GET("/products", productController.GetProducts)
	router.POST("/products", productController.CreateProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
	router.GET("/products/:id", productController.GetProductById)

	return router
}
