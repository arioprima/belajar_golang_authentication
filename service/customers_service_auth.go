package service

import (
	"context"
	"github.com/arioprima/belajar_golang_authentication/models/web"
)

type CustomersServiceAuth interface {
	Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error)
}
