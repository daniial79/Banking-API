package core

// Customer Core Object
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// Customer Secondary Port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

// Customer Primary Port
type CustomerService interface {
	GetAllCustomers() ([]Customer, error)
}
