package main

import (
	"fmt"
)

type Employee struct {
	ID     int
	name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

func (m *Manager) AddEmp(e Employee) {
	m.Employees = append(m.Employees, e)
}
func (m *Manager) RemoveEmp(idd int) {

	for i, emp := range m.Employees {
		if emp.ID == idd {
			m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
			return
		}
	}
}

func (m *Manager) GetSalId(idd int) float64 {

	var format float64
	for _, emp := range m.Employees {
		if emp.ID == idd {
			format = emp.Salary
		}
	}
	return format
}

func (m *Manager)Updt(e Employee , idd int) bool{

	for _ , emp := range m.Employees{
		if m.emp[i].ID == idd{
			 m.emp[i].ID = e.ID , 
			 m.emp[i].Age = e.Age ,
			float64( m.emp[i].Salary) = float64(e.Salary) ,
			 m.emp[i].name = e.name,
		}
	}
	return true

}

func main() {
	man := Manager{}

	man.AddEmp(Employee{ID: 1, name: "Chandra", Salary: 12000, Age: 12})
	man.AddEmp(Employee{ID: 2, name: "ds'lklmdra", Salary: 34000, Age: 18})
	man.AddEmp(Employee{ID: 3, name: "as.ndra", Salary: 23000, Age: 32})
	man.AddEmp(Employee{ID: 4, name: "adslkmdra", Salary: 50000, Age: 42})

	sla := man.GetSalId(2)

	fmt.Println(sla)

	updated := man.Updt(
		Employee{Name: "Updated Employee", Age: 26, Salary: 40000},2,)

	fmt.Println(updated)

}
