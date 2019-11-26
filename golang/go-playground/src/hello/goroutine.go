package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	// example1()

	// example2()

	// example3()

	// example3_2()
	example3_3()
}

func example1() {
	fmt.Println("example1 가용한 CPU 개수 :", runtime.NumCPU(), " / golang 실행시의 설정한 CPU 개수 :", runtime.GOMAXPROCS(0))

	go func() {
		fmt.Println("Hello, World")
	}()

	fmt.Scanln() // Hello world가 출력되기 전에 main 함수가 종료되어 버림.. 대기하기 위해 걸어놓은 함수
}

func example2() {
	for i := 0; i < 100; i++ {
		go func(n int) {
			r := rand.Intn(100)
			time.Sleep(time.Duration(r))

			fmt.Println(n, " sleep timer :", r)
		}(i)
	}

	fmt.Scanln()
}

func example3() {
	/*
	  runtime.GOMAXPROCS(runtime.NumCPU())
	  fmt.Println("가용한 CPU 개수 :", runtime.NumCPU(), " / golang 실행시의 설정한 CPU 개수 :", runtime.GOMAXPROCS(0))
	*/
	s := "hello world"

	for i := 0; i < 100; i++ {
		go func(n int) { // 해당 구문을 내 로컬 기준, 4개의 코어를 가지고 병렬로 실행한다.  순서가 뒤엉키게 출력된다.
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}

func example3_2() {
	runtime.GOMAXPROCS(1) // runtime.NumCPU())

	fmt.Println("example3_2 가용한 CPU 개수 :", runtime.NumCPU(), " / golang 실행시의 설정한 CPU 개수 :", runtime.GOMAXPROCS(0))

	s := "hello world"
	for i := 0; i < 100; i++ {
		go func(n int) { // 해당 구문을 1개의 코어로 실행하기에 순서대로 출력된다.
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}

func example3_3() {
	runtime.GOMAXPROCS(1)

	fmt.Println("example3_2 가용한 CPU 개수 :", runtime.NumCPU(), " / golang 실행시의 설정한 CPU 개수 :", runtime.GOMAXPROCS(0))

	s := "hello world"
	for i := 0; i < 100; i++ {
		go func() { // 변수 i를 클로저의 매개변수로 넘기지 않는다면 고루틴은 반복문이 끝난다음 생성되므로 항상 100이 출력된다.
			fmt.Println(s, i)
		}()
	}

	fmt.Scanln()
}
