package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

type address struct {
	person
	fName  string
	street string
}

// function receiver example - it will attach the method to receiver variable type
func (p person) fullName() string {
	return p.fName + " " + p.lName
}

func (a address) fullName() string {
	return a.person.fName + " " + a.person.lName + " " + a.fName + " " + a.street
}

func main() {
	p1 := person{"Gaulla", "Rajeswar Reddy", 37}
	a1 := address{
		person: person{"Gaulla",
			"Rajeswar Reddy",
			37,
		},
		fName:  "Janapriya",
		street: "attapur",
	}
	fmt.Println(p1.age)
	fmt.Println(a1.fName)
	fmt.Println(a1.person.fName)
	fmt.Println(p1.fullName())
	fmt.Println(a1.fullName())
	fmt.Println(a1.person.fullName())
	fmt.Println(p1.fullName() == a1.person.fullName())
}
