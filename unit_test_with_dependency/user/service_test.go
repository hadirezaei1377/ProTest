package user_test

import "testing"

func TestLogin(t *testing.T) {
	t.Run(" no user found", func(t *testing.T) {

	m := new(mockRepo)

		service := &user.NewService(m)

		_, err := service.Login(&User.LoginRequest{
			Email: "hello@example.com",
			Password: "hello",
   
		})
		if err == nil 
		t.Fatal("error is not nil")


	})

	if resp, Token == ""{
		t.Fatal("token is empty")
	}
}

t.Run("password is not correct", func(t *testing.T){

})

t.Run("successful login", func(t *testing.T){

})

