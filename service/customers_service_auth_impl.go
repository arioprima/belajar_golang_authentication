package service

import (
	"context"
	"database/sql"
	"github.com/arioprima/belajar_golang_authentication/config"
	"github.com/arioprima/belajar_golang_authentication/helper"
	"github.com/arioprima/belajar_golang_authentication/models/web"
	"github.com/arioprima/belajar_golang_authentication/repository"
	"github.com/arioprima/belajar_golang_authentication/utils"
	"github.com/go-playground/validator/v10"
)

type CustomersServiceAuthImpl struct {
	CustomersAuthRepository repository.CustomersRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewCustomersServiceAuth(customersAuthRepository repository.CustomersRepository, DB *sql.DB, validate *validator.Validate) CustomersServiceAuth {
	return &CustomersServiceAuthImpl{CustomersAuthRepository: customersAuthRepository, DB: DB, Validate: validate}
}

func (service *CustomersServiceAuthImpl) Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error) {
	//TODO implement me
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.Defer(tx)

	customers, err := service.CustomersAuthRepository.FindByUsername(ctx, tx, request.Username)
	helper.PanicIfError(err)

	config, _ := config.LoadConfig()
	verify_error := utils.CheckPasswordHash(customers.Password, request.Password)
	if verify_error != false {
		panic(verify_error)
	}

	token, err_token := utils.GenerateToken(config.TokenExpiresIn, customers.Id, config.TokenSecret)
	if err_token != nil {
		panic(err_token)
	}

	return web.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}, nil
}
