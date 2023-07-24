package repository

import (
	"context"
	"database/sql"
	"github.com/arioprima/belajar_golang_authentication/models/entity"
)

type CustomersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, customers entity.Customer) entity.Customer
	Update(ctx context.Context, tx *sql.Tx, customers entity.Customer) entity.Customer
	Delete(ctx context.Context, tx *sql.Tx, customers entity.Customer)
	FindById(ctx context.Context, tx *sql.Tx, customersId string) (entity.Customer, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.Customer, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Customer
}
