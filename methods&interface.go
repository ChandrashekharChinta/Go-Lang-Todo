1. Basic Method

func (receiver ReceiverType) MethodName() {
// method body
}

Example:

func (m Manager) Display() {
fmt.Println("Manager")
}

2. Method with Pointer Receiver

func (receiver *ReceiverType) MethodName() {
}

Example:

func (m *Manager) UpdateName() {
m.Name = "John"
}

3. Method with One Parameter

func (receiver ReceiverType) MethodName(param Type) {
}

Example:

func (m *Manager) AddEmployee(name string) {
fmt.Println(name)
}

4. Method with Multiple Parameters

func (receiver ReceiverType) MethodName(param1 Type1, param2 Type2) {
}

Example:

func (m *Manager) AddEmployee(name string, age int) {
fmt.Println(name, age)
}

5. Method Returning One Value

func (receiver ReceiverType) MethodName() ReturnType {
return value
}

Example:

func (m *Manager) GetName() string {
return m.Name
}

6. Method Returning Multiple Values

func (receiver ReceiverType) MethodName() (Type1, Type2) {
return value1, value2
}

Example:

func (m *Manager) GetDetails()*(string, int) {
return m.Name, 25*}
``*

Usage:

name, age := m.Get*etails()

7. Method Returning Value + Error (Most Common)

func (r*ceiver ReceiverType) MethodName() *Type, error) {
*return value, nil
}

Example:

func *m *Manager) GetSalary() (float*4, error) {
return 50000,*nil
}


//////////////////////////////////---------INTERFACES-----------------///////////////////////////

To declare an interface, you only write the method signatures (no func keyword, no method body).

---------Struct

type Manager struct {
Name string
}

------Methods

func (m *Manager) AddEmployee(name string) {
fmt.Println(name)
}

func (m *Manager) GetName() string {
return m.Name
}

func (m *Manager) GetDetails() (string, int) {
return m.Name, 25
}

func (m *Manager) GetSalary() (float64, error) {
return 50000, nil
}


---------Interface

type ManagerService interface {
AddEmployee(name string)
GetName() string
GetDetails() (string, int)
GetSalary() (float64, error)
}

-------------------------------------------------------------------------------


Method												Interface Signature

func (m *Manager) AddEmployee(name string)			AddEmployee(name string)
func (m *Manager) GetName() string					GetName() string
func (m *Manager) GetDetails() (string, int)		GetDetails() (string, int)
func (m *Manager) GetSalary() (float64, error)		GetSalary() (float64, error)


-----------------------------------------------------------------------------------
package main

import "fmt"

type Manager struct {
	Name string
}

type ManagerService interface {
	AddEmployee(name string)
	GetName() string
	GetDetails() (string, int)
	GetSalary() (float64, error)
}

func (m *Manager) AddEmployee(name string) {
	fmt.Println("Added:", name)
}

func (m *Manager) GetName() string {
	return m.Name
}

func (m *Manager) GetDetails() (string, int) {
	return m.Name, 25
}

func (m *Manager) GetSalary() (float64, error) {
	return 50000, nil
}

func main() {
	var service ManagerService = &Manager{Name: "John"}
	//Create a Manager object named "John" and assign it to the service interface variable 
	// because *Manager implements all methods of ManagerService.

	service.AddEmployee("Alice")

	name := service.GetName()
	fmt.Println(name)

	//OUTPUT -- Added: Alice
	//			John
}