package controller

import (
	"net/http"
	"petshop/data/request"
	"petshop/data/response"
	"petshop/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) *ProductController{
	return &ProductController{productService: service}
}

func (controller *ProductController) FindAll(ctx *gin.Context){
	data, err := controller.productService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code: 500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code: 200,
		Status: "Ok",
		Data: data,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductController) FindById(ctx *gin.Context){
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code: 400,
			Message: err.Error(),
		})
		return
	}

	data, err := controller.productService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code: 404,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code: 200,
		Status: "Ok",
		Data: data,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductController) Create(ctx *gin.Context){
	req := request.CreateProductRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code: 400,
			Message: err.Error(),
		})
		return
	}

	err := controller.productService.Create(req)
	if err != nil {
		statusError := http.StatusInternalServerError
		if err.Error() == "product code already exists" {
			statusError = http.StatusConflict
		}
		
		ctx.JSON(statusError, response.ErrorResponse{
			Code: statusError,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code: 200,
		Status: "Ok",
		Data: req,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductController) Update(ctx *gin.Context){
	req := request.UpdateProductRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code: 400,
			Message: err.Error(),
		})
		return
	}

	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code: 404,
			Message: "Invalid id",
		})
		return
	}

	req.Id = id

	err = controller.productService.Update(req)
	if err != nil {
		statusError := http.StatusInternalServerError
		if err.Error() != "product code already exits" {
			statusError = http.StatusConflict
		}

		ctx.JSON(statusError, response.ErrorResponse{
			Code: statusError,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code: 200,
		Status: "Ok",
		Data: req,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductController) Delete(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code: 404,
			Message: "Invalid id",
		})
		return
	}

	_, err = controller.productService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code: 404,
			Message: "product not found",
		})
	}

	err = controller.productService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code: 400,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code: 200,
		Status: "Ok",
		Data: nil,
	}

	ctx.JSON(http.StatusOK, res)
}

