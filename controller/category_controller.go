package controller

import (
	"net/http"
	"petshop/data/request"
	"petshop/data/response"
	"petshop/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: service}
}

func (controller *CategoryController) FindAll(ctx *gin.Context) {
	data, err := controller.categoryService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *CategoryController) FindById(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Invalid ID",
		})
		return
	}

	data, err := controller.categoryService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *CategoryController) Create(ctx *gin.Context) {
	req := request.CreateCategoryRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	err := controller.categoryService.Create(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "category name already exists" {
			statusCode = http.StatusConflict
		}

		ctx.JSON(statusCode, response.ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   201,
		Status: "Created",
		Data:   req,
	}
	ctx.JSON(http.StatusCreated, res)
}

func (controller *CategoryController) Update(ctx *gin.Context) {
	req := request.UpdateCategoryRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Invalid ID",
		})
		return
	}

	req.Id = id

	err = controller.categoryService.Update(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "category name already exists" {
			statusCode = http.StatusConflict
		}

		ctx.JSON(statusCode, response.ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *CategoryController) Delete(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Invalid ID",
		})
		return
	}

	_, err = controller.categoryService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: "Category not found",
		})
		return
	}

	err = controller.categoryService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, res)
}
