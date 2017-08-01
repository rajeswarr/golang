package main

import "fmt"

/* to interface u can pass any type argument which implements the interface method.
   In the below example to result method we can pass Square type variable because Square
   implemented Shape interface all the methods. */

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) triple() float64 {
	return s.side * s.side * s.side
}

type Shape interface {
	area() float64
	triple() float64
}

func result(s Shape) {
	fmt.Println(s.area())
	fmt.Println(s.triple())
}

func main() {
	s := Square{10}
	result(s)
}
