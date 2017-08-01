package main

import "fmt"

func main() {

	tx := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		ts := make([]int, 0)
		for j := 0; j < 4; j++ {
			ts = append(ts, j)
		}
		tx = append(tx, ts)
	}
	fmt.Println(tx)
}
