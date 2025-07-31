package database

import "github.com/sancheschris/goexpert/9-APIS/internal/entity"

type UserInterface interface{
	Create(user *entity.User) error
	FindByEmail(emailId string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(product *entity.Product) error
}