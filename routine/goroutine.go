package routine

import "github.com/google/uuid"

type ComponyInterface interface {
	saveToDatabse()
	getType() ComponyInterface
}

type ComponyEnum string

const (
	ComponyNew       ComponyEnum = "Compony"
	VendorComponyNew ComponyEnum = "Vendor"
	EmployeeNew      ComponyEnum = "Employee"
)

type VendorCompony struct {
	Compony
	AccountNumber int
}

type Compony struct {
	id             string
	name           string
	founder        string
	avg_make_money float64
	country        []Country
}

type Country struct {
	id   string
	name string
}

type Employee struct {
	Compony
	Salary      float64
	first_name  string
	second_name string
}

func newVendor(
	name_of_compony string,
	founder string,
	avg_make_money float64,
	my_country Country,
	AccountNumber int,
) VendorCompony {

	my_comp := newCompony(
		name_of_compony,
		founder,
		avg_make_money, my_country,
	)
	vender := VendorCompony{
		my_comp,
		AccountNumber,
	}

	return vender
}

func newEmployee(name_of_compony string,
	founder string,
	avg_make_money float64,
	my_country Country,
	fNname string, sName string, salary float64,
) Employee {

	my_comp := newCompony(
		name_of_compony,
		founder, avg_make_money, my_country,
	)

	return Employee{
		my_comp,
		salary,
		fNname,
		sName,
	}
}

func newCompony(name_of_compony string,
	founder string,
	avg_make_money float64,
	my_country Country) Compony {

	return Compony{
		id:             uuid.New().String(),
		name:           name_of_compony,
		founder:        founder,
		avg_make_money: avg_make_money,
		country:        []Country{my_country},
	}

}

// we have table of facts and измерения
