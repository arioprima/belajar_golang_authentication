package web

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=1,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}
