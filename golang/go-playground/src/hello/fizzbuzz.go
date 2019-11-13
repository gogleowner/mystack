package main

import . "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			Println("FizzBuzz")
		case i%3 == 0:
			Println("Fizz")
		case i%5 == 0:
			Println("Buzz")
		default:
			Println(i)
		}
	}
}
