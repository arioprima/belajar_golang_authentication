package repository

import (
	"context"
	"database/sql"
	"github.com/arioprima/belajar_golang_authentication/models/entity"
)

type CustomersRepositoryImpl struct {
	db *sql.DB
}

func NewCustomersRepository(db *sql.DB) CustomersRepository {
	return &CustomersRepositoryImpl{db}
}

func (c *CustomersRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customers entity.Customer) (entity.Customer, error) {
	//TODO implement me
	SQL := "insert into customers (id, firstname, lastname, username, email, password, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		customers.Id,
		customers.Firstname,
		customers.Lastname,
		customers.Username,
		customers.Email,
		customers.Password,
		customers.CreatedAt,
		customers.UpdatedAt,
	)
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (c *CustomersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customers entity.Customer) (entity.Customer, error) {
	//TODO implement me
	SQL := "update customers set firstname = ?, lastname = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		customers.Firstname,
		customers.Lastname,
		customers.UpdatedAt,
	)
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (c *CustomersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customers entity.Customer) {
	//TODO implement me
	SQL := "delete from customers where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customers.Id)
	if err != nil {
		panic(err)
	}
}

func (c *CustomersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customersId string) (entity.Customer, error) {
	//TODO implement me
	SQL := "select id, firstname, lastname, username, email, created_at, updated_at from customers where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customersId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := entity.Customer{}
	for rows.Next() {
		err := rows.Scan(
			&customers.Id,
			&customers.Firstname,
			&customers.Lastname,
			&customers.Username,
			&customers.Email,
			&customers.CreatedAt,
			&customers.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
	}
	return customers, nil
}

func (c *CustomersRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.Customer, error) {
	//TODO implement me
	SQL := "select id, firstname, lastname, username, email, created_at, updated_at from customers where username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := entity.Customer{}
	for rows.Next() {
		err := rows.Scan(
			&customers.Id,
			&customers.Firstname,
			&customers.Lastname,
			&customers.Username,
			&customers.Email,
			&customers.CreatedAt,
			&customers.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
	}
	return customers, nil
}

func (c *CustomersRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Customer {
	//TODO implement me
	SQL := "select id, firstname, lastname, username, email, created_at, updated_at from customers"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
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
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}
	return customers
}
