package main

import "fmt"

/*func main() {

	a := totalValue(10, 20, 30, 40, 50)
	fmt.Println(a)
}*/

/*
func main() {
	values := []float64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	a := totalValue(values...)
	fmt.Println(a)
}
*/

func main() {
	values := []float64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	a := totalValue(values)
	fmt.Println(a)
}

/*
func totalValue(tt ...float64) float64 {
	fmt.Printf("%T \n", tt)
	total := 0.0
	for j, i := range tt {
		fmt.Println(j)
		fmt.Println(i)
		total = total + i
	}
	return total
}
*/

func totalValue(tt []float64) float64 {
	fmt.Printf("%T \n", tt)
	total := 0.0
	for j, i := range tt {
		fmt.Println(j)
		fmt.Println(i)
		total = total + i
	}
	return total
}
