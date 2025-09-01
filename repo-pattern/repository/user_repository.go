package repository

import (
	"fmt"
	"repo-user/entity"
)

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id int) (*entity.User, error)
	Update(id int) (*entity.User, error)
}

type InMemoryUserRepo struct {
	Users map[int]entity.User
}

func (r *InMemoryUserRepo) Save(user *entity.User) error {
	r.Users[user.ID] = *user
	return nil
}

func (r *InMemoryUserRepo) FindByID(id int) (*entity.User, error) {
	user, ok := r.Users[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &user, nil
}

func (r *InMemoryUserRepo) Update(user *entity.User) (*entity.User, error) {
	_, ok := r.Users[user.ID]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	r.Users[user.ID] = *user
	return user, nil
}