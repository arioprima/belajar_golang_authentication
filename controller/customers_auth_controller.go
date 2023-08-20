package controller

import (
	"github.com/arioprima/belajar_golang_authentication/helper"
	"github.com/arioprima/belajar_golang_authentication/models/web"
	"github.com/arioprima/belajar_golang_authentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomersAuthController struct {
	customersAuthController service.CustomersServiceAuth
}

func NewCustomersAuthController(customersAuthController service.CustomersServiceAuth) *CustomersAuthController {
	return &CustomersAuthController{customersAuthController: customersAuthController}
}

func (controller *CustomersAuthController) Login(ctx *gin.Context) {
	var request web.LoginRequest
	err := ctx.Bind(&request)
	helper.PanicIfError(err)

	token, err_token := controller.customersAuthController.Login(ctx, request)
	if err_token != nil {
		webResponse := web.Response{
			Status: "Bad Request",
			Data:   nil,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := web.Response{
		Status: "OK",
		Data:   token,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
