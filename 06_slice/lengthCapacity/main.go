package main

import "fmt"

func main() {
	/*
		s := []int{2, 3, 5, 7, 11, 13}

		s = s[1:4]
		fmt.Println(s)

		s = s[:2]
		fmt.Println(s)

		s = s[1:]
		fmt.Println(s)
	*/
	s := []int{2, 3, 5, 7, 11, 13}
	//printSlice(s)
	fmt.Println(s)
	// Slice the slice to give it zero length.
	//s = s[:0]
	s = s[1:4]
	//printSlice(s)
	fmt.Println(s)
	// Extend its length.
	s = s[2:4]
	//printSlice(s)
	fmt.Println(s)
	// Drop its first two values.
	s = s[1:]
	//printSlice(s)
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
