package controllers

import (
	"api/models"
	"api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type produtController struct {
	productUsecase usecase.ProductUsecase
}

func NewProdutController(usecase usecase.ProductUsecase) produtController {
	return produtController{
		productUsecase: usecase,
	}
}

func (pc *produtController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *produtController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	p, err := pc.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, p)
}
