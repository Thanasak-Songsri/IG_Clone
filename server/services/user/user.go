package user

import "github.com/gin-gonic/gin"

type UserResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	IsActive bool   `json:"isactive"`
}

type UserService interface {
	Getusers() ([]UserResponse, error)
	CreateUser(c *gin.Context) (bool, error)

	SignUp(c *gin.Context) (bool, error)
	SignIn(c *gin.Context) (bool, error)
	SignOut(c *gin.Context) (bool, error)
}
