package config

import (
	"database/sql"
	"fmt"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"containers-us-west-210.railway.app",
		"6300",
		"postgres",
		"L5jQK0d4GDtgUbto68yA",
		"railway",
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
