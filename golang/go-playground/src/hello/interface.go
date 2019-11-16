package main

import (
	. "fmt"
	"strconv"
)

type Hello interface {
}

type MyInt int

func (i MyInt) Print() {
	Println(i)
}

type Printer interface {
	Print()
}

type Rectangle struct {
	width, height int
}

func (rect Rectangle) Print() {
	Println("width :", rect.width, "/ height :", rect.height)
}

// 빈 인터페이스 사용
func f1(arg interface{}) {
	Println(arg)
}

type Any interface{}

func f2(arg Any) {
	Println(arg)
}

func formatString(arg interface{}) string {
	switch arg.(type) { // arg.(type) 문은 switch에서만 쓸 수 있다.
	case int:
		i := arg.(int)         // type assertion
		return strconv.Itoa(i) // int i를 문자열로 변환. 그 반대는 Atoi
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32) // f를 문자열로 변환
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)
	case string:
		s := arg.(string)
		return s
	default:
		return "Error"
	}
}

type Person struct {
	name string
	age  int
}

func formatStringWhenHasPerson(arg interface{}) string {
	var result string
	switch arg.(type) {
	case Person:
		p := arg.(Person)
		result = p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		result = p.name + " " + strconv.Itoa(p.age)
	default:
		result = "Error"
	}

	return result
}

func main() {
	var h Hello
	Println(h)

	var i MyInt = 5
	var p Printer
	p = i     // Printer <- MyInt
	p.Print() // MyInt.Print() 메소드가 실행됨.

	var p2 Printer
	rect := Rectangle{10, 20}
	p2 = rect
	p2.Print()

	// 선언과 동시에 초기화
	pArr := []Printer{
		MyInt(5),
		Rectangle{10, 20},
	}

	for _, value := range pArr {
		Print(value, " -> ")
		value.Print()
	}

	// 타입이 특정 인터페이스를 구현하는지 검사
	if v, ok := interface{}(p).(Printer); ok {
		Println("타입이 특정 인터페이스를 구현하는지 검사", v, ok)
	}

	f1(p)
	f2(p)

	Println(formatString(1), formatString(1.5), formatString("Hello"))

	person := Person{"lee", 10}
	Println(formatStringWhenHasPerson(person))
	Println(formatStringWhenHasPerson(&person))

	var any interface{}
	any = Person{"lim", 10}
	// any = Rectangle{10, 20} // -> will print "not person"
	if v, ok := any.(Person); ok {
		Println(v, ok)
	} else {
		Println("not person ", any)
	}

}
