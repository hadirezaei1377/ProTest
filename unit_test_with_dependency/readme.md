we have an API ... we should  get email and password from users and generate a JWT token.

at first we use database for checking, service cant act without database object, so we injected database as an object to service.

this is raw code :
package user

import "gorm.io/gorm"

type service struct {
	db *gorm.DB // this is a dependency for this service
}

type User struct {
	ID       uint
	Email    string
	Password string
	Name     string
}

type LoginRequest struct {
	Email    string
	Password string
}

// this is a token
type LoginResponse struct {
	Token string
}

func (s *service) Login(request *LoginRequest) (*LoginResponse, error) {
	// todo

	/*
		1-get the users from databases by email
		2-check the users password
		3-generate JWT token
		4-return the response

	*/

	// 1
	user := New(User) // check if this user is exist in database or not
	err := s.db.Model(User).Where("email = ?", request.Email.First(User).Error


if err != nil {
	return nil, fmt.Errorf("user not found")
})

// now we can check the password is correct or not
//2
if User.Password != request.Password{
	return nil, fmt.Errorf("password is not correct")
}

// if the password was correct ...
//3
token := "234587hafduig8gnfhskjsadjfgaskdfj"

//4
return &LoginResponse{Token: token, nil}
}


// reading data from database is a dependency







and the test is :
package user_test

import "testing"

func TestLogin(t *testing.T) {
	t.Run(" no user found", func(t *testing.T) {

		db := &gorm.DB{}

		service := &user.NewService(db)

		_, err := service.Login(&User.LoginRequest{
			Email: "hello@example.com",
			Password: "hello",

		})
		if err == nil 
		t.Fatal("error is not nil")

	})
}

t.Run("password is not correct", func(t *testing.T){

})

t.Run("successful login", func(t *testing.T){

})









...
we musnt in our service logic , test another logics and another dependencies
it means for example we shouldnt check whether the database has connection or not.
if we check this connection , its not unit test longer.
we must decouple service module from another modules.
so I will change all above code .