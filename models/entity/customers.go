package entity

import "time"

type Customer struct {
	Id        string
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
