package main

import (
	"github.com/arioprima/belajar_golang_authentication/config"
	"github.com/arioprima/belajar_golang_authentication/controller"
	"github.com/arioprima/belajar_golang_authentication/helper"
	"github.com/arioprima/belajar_golang_authentication/repository"
	"github.com/arioprima/belajar_golang_authentication/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	loadConfig, err := config.LoadConfig()
	helper.PanicIfError(err)
	db, err := config.ConnectionDB(&loadConfig)
	helper.PanicIfError(err)

	validate := validator.New()

	customersRespository := repository.NewCustomersRepositoryImpl(db)

	customersService := service.NewCustomersServiceImpl(customersRespository, db, validate)

	customersController := controller.NewCustomersControllerImpl(customersService)

	router := gin.Default()

	router.POST("/api/v1/customers", customersController.Create)

	router.Run(":8080")
}
