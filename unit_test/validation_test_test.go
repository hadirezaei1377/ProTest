package unit_test_test

import "github.com/gogo/protobuf/test"

// table test
func TestRegisterValidation(t *test.T) {

	type input struct {
		request *delivery.RegisterRequest
		result  bool
	}

	// all bevaiors that the function could has is here
	var testCases = []input{
		{request: &delivery.RegisterRequest{}, result: false},
		{request: &delivery.RegisterRequest{Email: "hello@example.com"}, result: false},
		{request: &delivery.RegisterRequest{Email: "hello@example.com", password: "hello"}, result: true},
		{request: &delivery.RegisterRequest{Email: "hello@example.com", password: "hello", TimeZone: "Tehran"}, result: false},
		{request: &delivery.RegisterRequest{Email: "hello@example.com", password: "hello", TimeZone: "Asia/Tehran"}, result: true},
	}

	validate := validation.Newinputvalidation()
	for _, tc := range testCases {
		result := validate.RegisterValidation(tc.request)
		if result != nil && tc.result {
			t.Fatalf("wants %v gets %v", tc.result, result)
		}
	}
}
