package controller

import (
	"net/http"
	"petshop/data/request"
	"petshop/data/response"
	"petshop/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BranchController struct {
	branchService service.BranchService
}

func (controller BranchController) FindAll(ctx *gin.Context) {
	data, err := controller.branchService.FindAll()

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

func (controller *BranchController) FindById(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)

	data, err := controller.branchService.FindById(id)

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

func (controller *BranchController) Create(ctx *gin.Context) {
	req := request.CreateBranchRequest{}
	ctx.ShouldBindBodyWithJSON(req)

	err := controller.branchService.Create(req)

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

func (controller *BranchController) Update(ctx *gin.Context){
	req := request.UpdateBranchRequest{}
	err := ctx.ShouldBindBodyWithJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code: 400,
			Message: err.Error(),
		})
	}

	branchId := ctx.Param("id")
	id, err := strconv.Atoi(branchId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code: 400,
			Message: "Invalid Id",
		})
		return
	}

	req.Id = id

	err = controller.branchService.Update(req)
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
		Data: req,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *BranchController) Delete(ctx *gin.Context){
	branchId := ctx.Param("id")
	id, err := strconv.Atoi(branchId)

	_, err = controller.branchService.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code: 404,
			Message: "Id Not Found",
		})
		return
	}

	err = controller.branchService.Delete(id)
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
		Data: nil,
	}

	ctx.JSON(http.StatusOK, res)
}
