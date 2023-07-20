package entity

import "time"

type Admin struct {
	Id        string
	email     string
	password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
