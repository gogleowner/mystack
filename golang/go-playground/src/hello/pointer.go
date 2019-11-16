package main

import . "fmt"

func main() {
	var numPtr *int = new(int)
	var num int = 1

	numPtr = &num    // 포인터형 변수 -> 메모리 주소만 저장
	Println(*numPtr) // 역참조 -> 메모리 주소로 접근하여 값을 가져옴
	Println(numPtr)  // 포인터 변수에 저장된 메모리 주소
	Println(&num)    // 변수 num의 주소

	// 메모리 주소를 직접 대입하거나 포인터 연산을 허용하지 않는다.
	/*
	  var numPtr2 *int = new(int)

	  numPtr2++ // compile error
	  numPtr = 0x00000001 // compile error

	  ./pointer.go:17:10: invalid operation: numPtr2++ (non-numeric type *int)
	  ./pointer.go:18:10: cannot use 1 (type int) as type *int in assignment
	*/

	// 함수에서 포인터형 매개변수 사용하기
	var n int = 1
	hello(n)
	Println(n) // 1 -> 바깥의 n 값에는 영향이 없다. 매개변수에서 일반 자료형을 사용하면 값이 복사

	helloWithPointer(&n) // 변수 n의 주소를 넘김
	Println(n)           // 2
}

func hello(n int) {
	n = 2
}

func helloWithPointer(n *int) {
	Println("in helloWithPointer", n, "value is ", *n) // 변수 n은 주소값.
	*n = 2
}
