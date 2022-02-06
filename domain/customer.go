package domain

import (
	"github.com/golang/Hexagonal-golangBancking/dto"
	"github.com/golang/Hexagonal-golangBancking/errs"
	//"google.golang.org/grpc/status"
)

type Customer struct {
	Id          string  `db:"customer_id"`//`json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zipcode"`
	DateofBirth string  `db:"date_of_birth"`//`json:"dateof_birth" xml:"dateof"`
	Status      string `json:"status" xml:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func(c Customer) ToDto() dto.CustomerResponse {
	
	return dto.CustomerResponse {
		Id: c.Id,
		Name: c.Name,
		City: c.City,
		Zipcode: c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status: c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
