package service

import (
	"errors"
	"log"
	"strconv"

	"github.com/Bash360/go-template/internal/modules/users/dto"
	"github.com/Bash360/go-template/internal/modules/users/model"
	"github.com/Bash360/go-template/internal/platform/repository"
)

type UserService struct {
	userRepo repository.IRepository[model.User]
}

func NewUserService(userRepo repository.IRepository[model.User]) *UserService {
	return &UserService{userRepo: userRepo}
}

func (self *UserService) AddUser(user dto.User) error {
	err := self.userRepo.Create(&model.User{Name: user.Name, Email: user.Email})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (self *UserService) GetUsers() []model.User {
	return self.userRepo.FindAll()
}

func (self *UserService) GetUserByID(id uint) (*model.User, error) {
	user := self.userRepo.FindById(id)
	if user == nil {
		return nil, errors.New(strconv.Itoa(int(id)) + " invalid id")
	}

	return user, nil
}
