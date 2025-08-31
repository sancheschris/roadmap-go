package entity


type User struct {
	ID int
	Name string
	Email string
	Age int
}

func NewUser(id int, name, email string, age int) User {
	return User{
		ID: id,
		Name: name,
		Email: email,
		Age: age,
	}
}