package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["p1"] = 3
	myMap["p2"] = 5

	fmt.Println(myMap)

	// checking for key existence

	v, ok := myMap["p3"]
	//fmt.Println("Is key exist : ", ok, v)
	if ok {
		fmt.Println("value :", v)
	} else {
		fmt.Println("key dont exist")
	}

	// deleting key from map

	delete(myMap, "p2")
	fmt.Println("myMap after deletion:", myMap)

	//initializing map

	myMap1 := map[string]int{"s1": 1, "s2": 4}

	fmt.Println(myMap1)

}
