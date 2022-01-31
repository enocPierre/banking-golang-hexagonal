package domain


type CustomerRepositoryStub struct {
	Customers []Customer
}

func (S CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return S.Customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"1001", "Pierre", "Sao Paulo", "1234567890", "2022-31-01", "1"},
		{"1002", "Pietra", "Sao Paulo", "1234567890", "2022-31-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}