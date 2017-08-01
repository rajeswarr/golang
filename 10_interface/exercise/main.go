package main

import (
	"fmt"
	"sort"
)

type SortInt []int

func (input SortInt) Len() int {
	return len(input)
}

func (input SortInt) Swap(i, j int) {
	input[i], input[j] = input[j], input[i]
}

func (input SortInt) Less(i, j int) bool {
	return input[i] > input[j]
}

type people []string

func (input people) Len() int {
	return len(input)
}

func (input people) Swap(i, j int) {
	input[i], input[j] = input[j], input[i]
}

func (input people) Less(i, j int) bool {
	return input[i] > input[j]
}

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	s := []string{"Zeno", "John", "Al", "Jenny"}
	s1 := []string{"Rajeswar", "Reddy", "Gaulla", "Vishwa"}

	fmt.Println(n)
	//sort.Sort(SortInt(n))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	/*  p := people{"Zeno", "John", "Al", "Jenny"}  and p := people(s1) are equal */

	//p := people{"Zeno", "John", "Al", "Jenny"}
	p := people(s1)
	fmt.Println(people(s1))
	fmt.Println(p) //or fmt.Println(people(s1))
	sort.Sort(p)   //  or sort.Sort(people(s1))
	fmt.Println(p) // or fmt.Println(people(s1))
}
