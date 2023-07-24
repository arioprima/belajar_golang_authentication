package web

type CustomersCreateRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}
