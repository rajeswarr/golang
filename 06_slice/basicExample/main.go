package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6}

	for i, value := range mySlice {
		fmt.Println(i, " - ", value)
	}

	fmt.Println(mySlice)

}
