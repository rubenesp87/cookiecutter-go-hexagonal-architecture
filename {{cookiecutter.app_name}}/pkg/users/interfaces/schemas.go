package interfaces

// RegisterUserRequest echo struct
type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Age      int    `json:"age"`
}

// LoginUserRequest echo struct
type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
