package user

import (
	"fmt"
)

// we have no relation with dependency behavior, but also we want to check tho log in, so we use mocking

type mockRepo struct{}

func (m *mockRepo) GetUserByEmail(email string) (*model.User, error) {
	// we set the dependency behavior as mock :
	if email == "" {
		return nil, fmt.Errorf("")
	}
	return &model.User{
		ID:       1,
		Email:    "email",
		password: "hello",
		Name:     "hello",
	}, nil
}

type Repository interface {
	GetUserByEmail(email string) (*model.User, error)
	// in this mode our dependency is not gorm , but also it is an interface
}

type Service struct {
	repo Repository
}

type LoginRequest struct {
	Email    string
	Password string
}

// this is a token
type LoginResponse struct {
	Token string
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Login(request *LoginRequest) (*LoginResponse, error) {
	// todo

	/*
		1-get the users from databases by email
		2-check the users password
		3-generate JWT token
		4-return the response

	*/

	user, err := s.repo.GetUserByEmail(request.Email)

	if err != nil {
		return nil, fmt.Errorf("usern not found")
	}

	// 1
	user := New(User) // check if this user is exist in database or not

	// now we can check the password is correct or not
	//2
	if User.Password != request.Password {
		return nil, fmt.Errorf("password is not correct")
	}

	// if the password was correct ...
	//3
	token := "234587hafduig8gnfhskjsadjfgaskdfj"

	//4
	return &LoginResponse{Token: token, nil}
}

// reading data from database is a dependency
