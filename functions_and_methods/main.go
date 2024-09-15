package main

import "fmt"

type rectange struct {
	length, breadth int
}

func calSum(num1, num2 int) int {
	return num1 + num2
}

func (r rectange) area() int {
	return r.length * r.breadth
}

func main() {
	total := calSum(10, 20)
	fmt.Printf("Total : %d\n", total)
	r := rectange{10, 20}
	fmt.Printf("Area : %d\n", r.area())
}
