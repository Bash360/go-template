package dto

type User struct {
	Name  string `json:"name" binding:"required,min=3,max=20"`
	Email string `json:"email" binding:"required,email"`
}
