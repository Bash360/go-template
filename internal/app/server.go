package app

import (
	"github.com/Bash360/go-template/internal/modules"
	UHandler "github.com/Bash360/go-template/internal/modules/users/handler"
	"github.com/Bash360/go-template/internal/modules/users/model"
	URepo "github.com/Bash360/go-template/internal/modules/users/repository"
	UService "github.com/Bash360/go-template/internal/modules/users/service"
	"github.com/Bash360/go-template/internal/platform/database"
	platform "github.com/Bash360/go-template/internal/platform/repository"
)

func RunServer() {

	DB := database.Connect()

	userBaseRepo := platform.NewBaseRepo[model.User](DB)

	userRepo := URepo.NewUserRepository(userBaseRepo)

	userService := UService.NewUserService(userRepo)

	userHandler := UHandler.NewUserHandler(userService)

	routerOptions := modules.RouterOptions{UserHandler: userHandler}

	server := modules.SetUpRouter(routerOptions)

	server.Run()
}
