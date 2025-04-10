package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Age      int8	`json:"age" binding:"required,min=16,max=100"`
}
