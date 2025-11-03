package handler

import (
	"net/http"
	"strconv"

	"github.com/Bash360/go-template/internal/modules/users/dto"
	"github.com/Bash360/go-template/internal/modules/users/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (self *UserHandler) GetUsers(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, self.userService.GetUsers())
}

func (self *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid user id type"})
		return
	}
	user, err := self.userService.GetUserByID(uint(intId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (self *UserHandler) AddUser(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := self.userService.AddUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
