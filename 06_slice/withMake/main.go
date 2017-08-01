package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("mySlice length :", len(mySlice))
	fmt.Println("mySlice capacity :", cap(mySlice))

	for i := 0; i < 81; i++ {
		mySlice = append(mySlice, i)

		fmt.Println("mySlice length :", len(mySlice), " mySlice capacity :", cap(mySlice), " value :", mySlice[i])
	}

}
