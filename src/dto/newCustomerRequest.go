package dto

import "github.com/daniial79/Banking-API/src/errs"

type NewCustomerRequest struct {
	Name        string `json:"name"`
	DateOfbirth string `json:"date_of_birth"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	Status      string `json:"status"`
}

func (req NewCustomerRequest) Validate() *errs.AppError {
	if req.Name == "" {
		return errs.NewValidationErr("Name must be provided")
	}

	if req.City == "" {
		return errs.NewValidationErr("City must be provided")
	}

	if len(req.DateOfbirth) != 10 {
		return errs.NewValidationErr("Date of birth must be provided with forma YYYY-MM-DD")
	}

	if req.Zipcode == "" {
		return errs.NewValidationErr("zipcode must be provided")
	}

	return nil
}
