package user

import (
	"igclone/models"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	UserCreate(c *gin.Context) (bool, error)
	ProfileCreate() (models.Userprofile, error)
	ProfileEdit(c *gin.Context) (bool, error)
}
