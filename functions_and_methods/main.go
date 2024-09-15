package main

import "fmt"

type rectangle struct {
	length, breadth int
}

func calSum(num1, num2 int) int {
	return num1 + num2
}

func (r rectangle) area() int {
	return r.length * r.breadth
}

func main() {
	var num1, num2, l, b int
	fmt.Scanf("%d %d %d %d", &num1, &num2, &l, &b)
	total := calSum(num1, num2)
	fmt.Printf("Total : %d\n", total)
	r := rectangle{l, b}
	fmt.Printf("Area : %d\n", r.area())
}
