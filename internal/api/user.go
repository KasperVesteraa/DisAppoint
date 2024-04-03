package api

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

func CreateUser(id, name, email, password string) (*User, error) {

	return &User{Id: id, Name: name, Email: email, Password: password}, nil
}
