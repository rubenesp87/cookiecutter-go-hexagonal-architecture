package interfaces

type CreateUserRequest struct {
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Age     int    `json:"age" validate:"required"`
}
