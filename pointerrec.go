package main

import (
	"fmt"
)

type Stud struct {
	name string
	age  int32
}

func (s *Stud) ChangeName(name string) {

	s.name = name
}

func main() {

	s := Stud{name: "Arjun", age: 12}

	fmt.Println(s.name)

	s.ChangeName("Bheem")

	fmt.Println(s.name)
}
