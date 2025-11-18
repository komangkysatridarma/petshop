package controller

import (
	"net/http"
	"petshop/data/request"
	"petshop/data/response"
	"petshop/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService service.RoleService
}

func NewRoleController(service service.RoleService) *RoleController {
	return &RoleController{roleService: service}
}

func (c *RoleController) FindAll(ctx *gin.Context) {
	data, err := c.roleService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	})
}

func (c *RoleController) FindById(ctx *gin.Context) {
	roleIdStr := ctx.Param("id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Invalid Role ID",
		})
		return
	}

	data, err := c.roleService.FindById(roleId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "Ok",
		Data:   data,
	})
}

func (c *RoleController) Create(ctx *gin.Context) {
	req := request.CreateRoleRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Invalid Request Body",
		})
		return
	}

	if err := c.roleService.Create(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:   201,
		Status: "Created",
		Data:   nil,
	})
}

func (c *RoleController) Update(ctx *gin.Context) {
	req := request.UpdateRoleRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Invalid Request Body",
		})
		return
	}

	roleIdStr := ctx.Param("id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Invalid ID format",
		})
		return
	}
	req.Id = roleId

	if err := c.roleService.Update(req); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "role not found" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, response.ErrorResponse{
			Code:    status,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "Ok",
		Data:   req,
	})
}

func (c *RoleController) Delete(ctx *gin.Context) {
	roleIdStr := ctx.Param("id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "Invalid Role ID",
		})
		return
	}

	if err := c.roleService.Delete(roleId); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "role not found" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, response.ErrorResponse{
			Code:    status,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	})
}