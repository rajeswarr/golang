package main

import "fmt"

func half(i int) (int, bool) {
	return i / 2, i%2 == 0
}

func main() {
	h, b := half(4)
	fmt.Println(h, b)
}
