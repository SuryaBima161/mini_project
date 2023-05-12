package payload

type (
	CreateUserRequest struct {
		Name     string `json:"name" form:"name" validate:"required,max=20"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	CreateUserResponse struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginUserRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	LoginUserResponse struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}
	GetUserByIdResponse struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UpdateUserRequest struct {
		Name     string `json:"name" form:"name" validate:"max=20"`
		Email    string `json:"email" form:"email" validate:"email"`
		Password string `json:"password" form:"password"`
	}
)
