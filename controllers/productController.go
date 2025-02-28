package controllers

import (
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
