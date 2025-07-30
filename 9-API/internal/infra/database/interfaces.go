package database

import "github.com/sancheschris/goexpert/9-APIS/internal/entity"

type UserInterface interface{
	Create(user *entity.User) error
	FindByEmail(emailId string) (*entity.User, error)
}