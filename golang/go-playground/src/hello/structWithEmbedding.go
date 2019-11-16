package main

import . "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	Println("Hello")
}

type StudentHasA struct {
	person Person // Has-a 관계 -> 학생이 사람을 가지고 있다.
	school string
	grade  int
}

type StudentIsA struct {
	Person // 필드명 없이 타입만 선언하면 is-a 관계. 구조체가 타입을 포함한다. -> 학생은 사람이다.
	school string
	grade  int
}

// greeting() 메소드를 오버라이딩하여 구현할 수 있다.
func (p *StudentIsA) greeting() {
	Println("Hello Students")
}

func main() {
	var s StudentHasA
	s.person.greeting()

	var s2 StudentIsA
	s2.Person.greeting()
	s2.greeting() // 바로 호출할 수 있다.
}
