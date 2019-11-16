package main

import . "fmt"

type Rectangle struct {
	width  int
	height int
	// width, height int // 한줄로 나열 가능
}

func main() {
	Println("struct 기본 선언")
	BasicDefinition()
	Println()
	Println("struct 생성자 패턴 활용")
	Constructor()
	Println()
	Println("함수에서 구조체 매개변수 사용")
	StructToMethodParameter()
}

func BasicDefinition() {
	var rect Rectangle // 기본 값으로 초기화됨.

	var rectPointer *Rectangle
	rectPointer = new(Rectangle)

	rectPointer2 := new(Rectangle)

	var rectWithValue Rectangle = Rectangle{10, 20}
	rectWithValue2 := Rectangle{10, 20}
	rectWithValue3 := Rectangle{width: 10, height: 20}

	Println(rect, rectPointer, rectPointer2, rectWithValue, rectWithValue2, rectWithValue3)

	// 값을 변경할 때
	rect.height = 20
	rectPointer.height = 62
	rectPointer2.height = 789

	Println(rect, rectPointer, rectPointer2)
}

func Constructor() {
	rect := NewRectangle(20, 10)
	Println(rect)

	rect2 := &Rectangle{width: 10}

	Println(rect, rect2)
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func StructToMethodParameter() {
	rect := Rectangle{30, 10}
	Println(RectangleArea(&rect)) // 구조체의 주소를 넘김. 일반 구조체를 넘기면 값이 복사된다.
}

func RectangleArea(rect *Rectangle) int {
	return rect.width + rect.height
}
