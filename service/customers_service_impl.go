package service

import (
	"context"
	"database/sql"
	"github.com/arioprima/belajar_golang_authentication/helper"
	"github.com/arioprima/belajar_golang_authentication/models/entity"
	"github.com/arioprima/belajar_golang_authentication/models/web"
	"github.com/arioprima/belajar_golang_authentication/repository"
	"github.com/arioprima/belajar_golang_authentication/utils"
	"github.com/go-playground/validator/v10"
	"time"
)

type CustomersServiceImpl struct {
	CustomersRepository repository.CustomersRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewCustomersServiceImpl(customersRepository repository.CustomersRepository, DB *sql.DB, validate *validator.Validate) CustomersService {
	return &CustomersServiceImpl{CustomersRepository: customersRepository, DB: DB, Validate: validate}
}

func (service *CustomersServiceImpl) Create(ctx context.Context, request web.CustomersCreateRequest) web.CustomersResponse {
	//TODO implement me
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.Defer(tx)

	paswordHash, err := utils.HashPassword(request.Password)
	helper.PanicIfError(err)

	customers := entity.Customer{
		Id:        utils.Uuid(),
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Username:  request.Username,
		Email:     request.Email,
		Password:  paswordHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	customers = service.CustomersRepository.Create(ctx, tx, customers)

	return utils.CustomersResponse(customers)
}

func (service *CustomersServiceImpl) Update(ctx context.Context, request web.UpdateCustomersRequest) web.CustomersResponse {
	//TODO implement me
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.Defer(tx)

	customers, err := service.CustomersRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	customers.Firstname = request.Firstname
	customers.Lastname = request.Lastname
	customers.UpdatedAt = time.Now()

	customers = service.CustomersRepository.Update(ctx, tx, customers)

	return utils.CustomersResponse(customers)
}

func (service *CustomersServiceImpl) Delete(ctx context.Context, customersId string) {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.Defer(tx)

	customers, err := service.CustomersRepository.FindById(ctx, tx, customersId)
	helper.PanicIfError(err)

	service.CustomersRepository.Delete(ctx, tx, customers)
}

func (service *CustomersServiceImpl) FindById(ctx context.Context, customersId string) web.CustomersResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.Defer(tx)

	customers, err := service.CustomersRepository.FindById(ctx, tx, customersId)
	helper.PanicIfError(err)

	return utils.CustomersResponse(customers)
}

func (service *CustomersServiceImpl) FindAll(ctx context.Context) []web.CustomersResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.Defer(tx)

	customers := service.CustomersRepository.FindAll(ctx, tx)

	var customersResponse []web.CustomersResponse
	for _, customer := range customers {
		customersResponse = append(customersResponse, utils.CustomersResponse(customer))
	}

	return customersResponse
}
