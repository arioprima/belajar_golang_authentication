package web

type UpdateCustomersRequest struct {
	Id        string `json:"id" validate:"required"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
}
