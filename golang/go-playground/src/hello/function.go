package example

import . "fmt"

func main() {
	hello()
	sumResult := sum(1, 2)

	Println(sumResult)
	Println(sum2(1, 3))

	// function result with multi result
	sumResult, difference := SumAndDiff(3, 1)
	Println(sumResult, difference)
	sumResult2, diff2 := SumAndDiff(3, 1)

	Println(sumResult2, diff2)

	// multi result
	a, _, c, _, e := helloMultiResult()
	Println(a, c, e)
	// variable argument
	Println(sumWithVariableArgument(1, 2, 3, 4, 5))
	Println(sumWithVariableArgument([]int{1, 2, 3, 4, 5}...))

  // recursive
	Println(factorial(5))

  // function assign to varible
  var sumFunc func(a int, b int) int = sum
  sumFunc2 := sum
  Println("sumFucn:", sumFunc(1, 2), "sumFunc2:", sumFunc2(1, 2))

  numberFunc := []func(int, int) int{sum, diff}

  Println(numberFunc)
  Println(numberFunc[0](1, 2), numberFunc[1](1, 2))

  numFuncMap := map[string]func(int, int) int {
    "sum" : sum,
    "diff" : diff,
  }

  Println(numFuncMap["sum"](1, 2), numFuncMap["diff"](1, 2))

  // anonymous function => usage closer, defer, goroutine
  func() {
    Println("Hello world")
  }()

  func(s string) {
    Println(s)
  }("Hello, world!")

  r := func(num1 int, num2 int) int {
    return num1 + num2
  }(1, 2)

  Println(r)

}

func hello() {
	Println("hello world")
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func diff(num1 int, num2 int) int {
	return num1 - num2
}


func sum2(num1 int, num2 int) (result int) {
	result = num1 + num2

	return
}

func SumAndDiff(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func SumAndDiff1(num1 int, num2 int) (sum int, diff int) {
	sum = num1 + num2
	diff = num1 - num2
	return
}

func helloMultiResult() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}

func sumWithVariableArgument(n ...int) int {
	total := 0

	for _, value := range n {
		total += value
	}

	return total
}

func factorial(n int) int {
	if n >= 1 {
		return n * factorial(n-1)
	} else {
		return 1
	}
}
