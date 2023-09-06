package unit_test

import (
	"fmt"
	"time"
)

// using test for this method
func (s *validate) RegisterValidation(req *delivery.RegisterRequest) error {
	const op = richerror.Operation("validation.RegisterValidation")

	err := validation.validateStruct(req,
		validation.Field(&req.Email, validation.Required),
		validation.Field(&req.Password, validation.Required))

	if err != nil {
		return richerror.New(err.Error()).withOperation(op)
	}

	if req.TimeZone != "" {
		if _, err := time.LoadLocation(req.TimeZone); err != nil {
			return richerror.New(fmt.Sprintf("#{req.TimeZone} is not valid time zone")).
				WithOperation(op).WithError(err)
		}
	}

	return nil
}

// the functions behavior:
// email be empty , password be empty , timezone be empty
