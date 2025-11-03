package repository

import (
	"github.com/Bash360/go-template/internal/modules/users/model"
	"github.com/Bash360/go-template/internal/platform/repository"
)

type UserRepository struct {
	*repository.Base[model.User]
}

func NewUserRepository(base *repository.Base[model.User]) *UserRepository {
	return &UserRepository{Base: base}
}
