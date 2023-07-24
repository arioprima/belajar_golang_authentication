package service

import (
	"context"
	"github.com/arioprima/belajar_golang_authentication/models/web"
)

type CustomersService interface {
	Create(ctx context.Context, request web.CustomersCreateRequest) web.CustomersResponse
	Update(ctx context.Context, request web.UpdateCustomersRequest) web.CustomersResponse
	Delete(ctx context.Context, customersId string)
	FindById(ctx context.Context, customersId string) web.CustomersResponse
	FindAll(ctx context.Context) []web.CustomersResponse
}
