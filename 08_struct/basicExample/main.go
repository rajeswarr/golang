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
}
