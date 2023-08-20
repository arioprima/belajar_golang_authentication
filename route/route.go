package route

import (
	"github.com/arioprima/belajar_golang_authentication/controller"
	"github.com/arioprima/belajar_golang_authentication/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(customersController *controller.CustomersControllerImpl, customersAuthController *controller.CustomersAuthController) *gin.Engine {
	service := gin.Default()
	middleware.SetupCorsMiddleware(service)
	router := service.Group("/api/v1/customers")

	router.POST("/register", customersController.Create)
	router.POST("/login", customersAuthController.Login)

	return service
}
