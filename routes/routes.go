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
	dbConnection := db.ConnectDB()

	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controllers.NewProdutoController(productUseCase)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
	})

	router.GET("/products", productController.GetProducts)
	router.POST("/products", productController.CreateProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
	router.GET("/products/:id", productController.GetProductById)

	return router
}
