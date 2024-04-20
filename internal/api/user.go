package api

import "github.com/google/uuid"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(id, name, email, password string) (*User, error) {

	return &User{Id: id, Name: name, Email: email, Password: password}, nil
}

func (u *User) CreateUuid() {
	u.Id = uuid.New().String()
}
