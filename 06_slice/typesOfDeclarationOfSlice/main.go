package main

import "fmt"

func main() {

	// 1 Method of slice declaration

	student := []string{}
	students := [][]string{}

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	// 2 Method of slice declaration

	var student1 []string
	var students1 [][]string

	fmt.Println(student1)
	fmt.Println(students1)
	fmt.Println(student1 == nil)

	// 3 Method of slice declaration -- it is prefered method

	student2 := make([]string, 10)
	students2 := make([][]string, 10)

	fmt.Println(student2)
	fmt.Println(students2)
	fmt.Println(student2 == nil)
}
