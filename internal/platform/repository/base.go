package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Base[T any] struct {
	DB *gorm.DB
}

func NewBaseRepo[T any](db *gorm.DB) *Base[T] {
	return &Base[T]{
		DB: db,
	}
}

func (self *Base[T]) Create(model *T) error {
	return self.DB.Create(model).Error
}

func (self *Base[T]) FindAll() []T {
	var models []T

	self.DB.Begin().Find(&models)
	return models
}

func (self *Base[T]) FindById(id uint) *T {
	var model *T

	self.DB.First(model, id)

	return model
}

func (self *Base[T]) Update(id uint, updates any) error {
	var model T
	result := self.DB.Model(&model).Where("id = ?", id).Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("record not found no row affected")
	}

	return nil
}

func (self *Base[T]) Delete(id uint) error {

	var model T

	result := self.DB.Delete(&model, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("record not found no row affected")
	}

	return nil

}
