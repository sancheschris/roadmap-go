package main

import (
	"fmt"
	"repo-user/entity"
	"repo-user/repository"
)



func main() {
	repo := repository.InMemoryUserRepo{
		Users: make(map[int]entity.User),
	}

	user := entity.NewUser(1, "Chris", "chris@gmail.com", 32)

	repo.Save(user)
	foundUser, err := repo.FindByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
}

