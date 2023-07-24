package utils

import (
	"github.com/arioprima/belajar_golang_authentication/models/entity"
	"github.com/arioprima/belajar_golang_authentication/models/web"
)

func CustomersResponse(customers entity.Customer) web.CustomersResponse {
	return web.CustomersResponse{
		Id:        customers.Id,
		Firstname: customers.Firstname,
		Lastname:  customers.Lastname,
		Username:  customers.Username,
		Email:     customers.Email,
		CreatedAt: customers.CreatedAt,
		UpdatedAt: customers.UpdatedAt,
	}
}
