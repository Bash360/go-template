package repository

type IRepository[T any] interface {
	FindById(id uint) *T
	FindAll() []T
	Create(entity *T) error
	Update(id uint, update any) error
	Delete(id uint) error
}
