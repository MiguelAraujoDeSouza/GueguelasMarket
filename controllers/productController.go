package controllers

import (
	"api/models"
	"api/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type produtController struct {
	productUsecase usecase.ProductUseCase
}

func NewProdutoController(usecase usecase.ProductUseCase) produtController {
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

func (pc *produtController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "ID is required")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "ID must be a number")
		return
	}

	product, err := pc.productUsecase.GetProductById(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *produtController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "ID is required")
		return
	}

	idInt, err := strconv.Atoi(id)

	err = pc.productUsecase.DeleteProduct(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
