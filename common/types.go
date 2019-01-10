package common

// RegisterForm is form binding struct for register
type RegisterForm struct {
	Name     string `json:"name" binding:"required,min=6,max=16"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=16"`
	Avatar   string `json:"avatar" binding:"omitempty,url"`
}

// RegisterForm is form binding struct for login
type LoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

type PostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CommentInput struct {
	Content string `json:"content" binding:"required"`
	PostID  string `json:"post_id" binding:"required"`
}
