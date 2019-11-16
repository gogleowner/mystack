package main

import . "fmt"

type Rectangle struct {
	width, height int
}

func (rect *Rectangle) area() int { // 리시버 벼수 정의(연결할 구조체 지정)
	return rect.width * rect.height
}

func (rect *Rectangle) scaleA(factor int) {
	rect.width *= factor
	rect.height *= factor

	Println("scaleA with pointer receiver variable :", rect)
}

func (rect Rectangle) scaleB(factor int) {
	rect.width *= factor
	rect.height *= factor

	Println("scaleB with pointer receiver variable :", rect)
}

func main() {
	Println("함수 정의시 리시버 변수를 통해 구조체의 값에 접근할 수 있다.수")
	rect := Rectangle{10, 20}
	Println(rect.area())

	Println("리시버 변수를 포인터로 받으면 원래 값이 변경되지만 값으로 받으면 원래의 값에는 영향을 미치지 않는다.")

	rect.scaleA(10)
	Println("after scaleA :", rect)
	rect2 := Rectangle{10, 20}
	rect2.scaleB(100)
	Println("after scaleB :", rect2)
}
