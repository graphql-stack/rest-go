package common

// RegisterForm is form binding struct for register
type RegisterForm struct {
	Username string `json:"username" binding:"omitempty,min=6,max=16"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

// RegisterForm is form binding struct for login
type LoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
