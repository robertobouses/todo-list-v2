package tasks

import (
	"fmt"
)

type CreateUserParams struct {
	email    string
	password string
}

func (t taskAppService) CreateUser(params CreateUserParams) (User, error) {
	if params.email == "" || params.password == "" {
		return User{}, fmt.Errorf("email y password no pueden estar vacíos")
	}

	return User{
		email:    params.email,
		password: params.password,
	}, nil
}
