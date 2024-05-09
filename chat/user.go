package chat

import (
	"errors"
	"fmt"
)

type User struct {
	Id       int
	Username string
}

var Users []User

func CreateUser(username string) (User, error) {

	//validar si el usuario ya existe
	for _, u := range Users {
		if u.Username == username {
			fmt.Println("El usuario ya existe")
			return User{}, errors.New("el usuario ya existe")
		}
	}

	u := User{Id: len(Users) + 1, Username: username}
	Users = append(Users, u)

	//AddUserToRoom(1, &u)

	fmt.Println(Users)
	return u, nil
}
