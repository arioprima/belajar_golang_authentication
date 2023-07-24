package repository

import (
	"context"
	"database/sql"
	"github.com/arioprima/belajar_golang_authentication/helper"
	"github.com/arioprima/belajar_golang_authentication/models/entity"
)

type CustomersRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomersRepositoryImpl(db *sql.DB) CustomersRepository {
	return &CustomersRepositoryImpl{DB: db}
}

func (c *CustomersRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customers entity.Customer) entity.Customer {
	//TODO implement me
	//pada kali ini akan menggunkan sintask sql query postgresql

	SQL := "INSERT INTO customers (id, firstname, lastname, username, email, password) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		customers.Id,
		customers.Firstname,
		customers.Lastname,
		customers.Username,
		customers.Email,
		customers.Password,
	)
	helper.PanicIfError(err)

	return customers
}

func (c *CustomersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customers entity.Customer) entity.Customer {
	//TODO implement me
	SQL := "UPDATE customers SET firstname = $1, lastname = $2, updated_at = $3 WHERE id = $4"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		customers.Firstname,
		customers.Lastname,
		customers.UpdatedAt,
		customers.Id,
	)
	helper.PanicIfError(err)

	return customers
}

func (c *CustomersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customers entity.Customer) {
	//TODO implement me
	SQL := "DELETE FROM customers WHERE id = $1"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		customers.Id,
	)
	helper.PanicIfError(err)
}

func (c *CustomersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customersId string) (entity.Customer, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username, email, created_at, updated_at FROM customers WHERE id = $1"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		customersId,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	customers := entity.Customer{}
	if rows.Next() {
		err := rows.Scan(
			&customers.Id,
			&customers.Firstname,
			&customers.Lastname,
			&customers.Username,
			&customers.Email,
			&customers.CreatedAt,
			&customers.UpdatedAt,
		)
		helper.PanicIfError(err)
		return customers, nil
	} else {
		return customers, err
	}
}

func (c *CustomersRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.Customer, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username, email, created_at, updated_at FROM customers WHERE username = $1"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		username,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	customers := entity.Customer{}
	if rows.Next() {
		err := rows.Scan(
			&customers.Id,
			&customers.Firstname,
			&customers.Lastname,
			&customers.Username,
			&customers.Email,
			&customers.CreatedAt,
			&customers.UpdatedAt,
		)
		helper.PanicIfError(err)
		return customers, nil
	} else {
		return customers, err
	}
}

func (c *CustomersRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Customer {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username, email, created_at, updated_at FROM customers"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(
			&customer.Id,
			&customer.Firstname,
			&customer.Lastname,
			&customer.Username,
			&customer.Email,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}
	return customers
}
