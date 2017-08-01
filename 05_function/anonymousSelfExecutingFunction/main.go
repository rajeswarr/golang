package main

import "fmt"

func main() {
	func() {
		i := 0
		i = i + 3
		fmt.Println(i)
	}()
}
