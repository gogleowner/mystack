package main

import . "fmt"

func main() {
	sum := func(a int, b int) int {
		return a + b
	}

	result := sum(1, 2)

	Println(result)

	num1, num2 := 3, 5
	c1 := func(x int) int {
		return num1*x + num2
	}

	Println(20 == c1(5))

	f := calc()

	Println(8 == f(1), 11 == f(2))
}

func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b
	}
}
