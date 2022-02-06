package domain

import (
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

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
