package main

import (
	"fmt"
	"log"
	"time"
)

/*
package fmt
- func Errorf(format string, a ...interface{}) error : 형식을 지정하여 error값을 만듦
package log
- func Fatal(v ...interface{}) : 에러 로그를 출력하고 프로그램을 완전히 종료
- func Panic(v ...interface{}) : 시간과 에러 문자열을 출력한 뒤 패닉을 발생시킴
- func Print(v ...interface{}) : 시간과 에러 메시지를 출력하며 프로그램을 종료하지 않음
 */
func main() {
	//example_1_fatal()
	//example_1_panic()
	//example_1_panicRecover()
	//example_1_logPrint()
	example_2_errorType()
}

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	} else {
		return "", fmt.Errorf("%d는 1이 아닙니다.", n)
	}
}

// 에러 문자열을 출력하고 프로그램을 완전히 종료한다.
func example_1_fatal() {
	s, e := helloOne(1)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(s)

	s2, e2 := helloOne(2)
	if e2 != nil {
		log.Fatal(e2)
	}
	fmt.Println(s2)
	fmt.Println("Hello world!")
}

// 런타임 중 패닉이 발생하여 콜스택이 출력된다.
func example_1_panic() {
	s, e := helloOne(1)
	if e != nil {
		log.Panic(e)
	}

	fmt.Println(s)

	s2, e2 := helloOne(2)
	if e2 != nil {
		log.Panic(e2)
	}
	fmt.Println(s2)
	fmt.Println("Hello world!")
}

func example_1_panicRecover() {
	defer func() {
		s := recover()
		fmt.Println("[recover]", s)
	}()

	s, e := helloOne(1)
	if e != nil {
		log.Panic(e)
	}

	fmt.Println(s)

	s2, e2 := helloOne(2)
	if e2 != nil {
		log.Panic(e2)
	}
	fmt.Println(s2)
	fmt.Println("Hello world!")
}

func example_1_logPrint() {
	s, e := helloOne(1)
	if e != nil {
		log.Print(e)
	}

	fmt.Println(s)

	s2, e2 := helloOne(2)
	if e2 != nil {
		log.Print(e2)
	}
	fmt.Println(s2)
	fmt.Println("Hello world!")
}


type HelloOneError struct {
	time time.Time
	value int
}

func (e HelloOneError) Error() string{
	return fmt.Sprintf("%v: %d는 1이 아닙니다.", e.time, e.value)
}

func helloOneWithType(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	} else {
		return "", HelloOneError{time.Now(), n}
	}
}

func example_2_errorType() {
	s, e := helloOneWithType(1)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(s)

	s2, e2 := helloOneWithType(2)
	if e2 != nil {
		log.Fatal(e2)
	}
	fmt.Println(s2)
	fmt.Println("Hello world!")

}

