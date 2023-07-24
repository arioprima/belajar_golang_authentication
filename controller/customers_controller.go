package controller

import (
	"github.com/arioprima/belajar_golang_authentication/models/web"
	"github.com/arioprima/belajar_golang_authentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomersControllerImpl struct {
	CustomersService service.CustomersService
}

func NewCustomersControllerImpl(customersService service.CustomersService) *CustomersControllerImpl {
	return &CustomersControllerImpl{CustomersService: customersService}
}

func (controller *CustomersControllerImpl) Create(ctx *gin.Context) {
	//TODO implement me
	customersCreateRequest := web.CustomersCreateRequest{}
	err := ctx.Bind(&customersCreateRequest)
	if err != nil {
		return
	}

	response := controller.CustomersService.Create(ctx, customersCreateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
