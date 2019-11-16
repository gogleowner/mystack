package main

import . "fmt"

// 만약 어떤 새가 오리처럼 걷고, 헤엄치고, 꽥꽥거리는 소리를 낸다면 나는 그 새를 오리라고 부를 것이다.
type Duck struct {
}

func (d Duck) quack() {
	Println("꽥꽥")
}
func (d Duck) feathers() {
	Println("오리에게 흰색, 회색 깃털이 있습니다")
}

type Person struct {
}

func (p Person) quack() {
	Println("이 사람이 오리를 흉내내네요.")
}
func (p Person) feathers() {
	Println("사람은 바닥에서 깃털을 주어서 보여 줍니다.")
}

type Quacker interface {
	quack()
	feathers()
}

func InTheForest(quacker Quacker) {
	quacker.quack()
	quacker.feathers()
}

func main() {
	var donald Duck
	var john Person

	quackers := []Quacker{
		donald,
		john,
	}

	for _, value := range quackers {
		InTheForest(value)
	}
}
