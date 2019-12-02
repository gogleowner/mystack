package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	// example_0_printType()
	// example_1_structTag()
	example_2_pointerAndInterface()
}

/*
reflection은 실행시점에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는 기능
*/
func example_0_printType() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num))

	var num32 int32 = 1
	fmt.Println(reflect.TypeOf(num32))

	var num64 int64 = 1
	fmt.Println(reflect.TypeOf(num64))

	var fNum float32 = 1.3
	fmt.Println(reflect.TypeOf(fNum))

	fNum2 := 1.3
	fmt.Println(reflect.TypeOf(fNum2))

	var fNum64 float64 = 1.3
	fmt.Println(reflect.TypeOf(fNum64))

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data)) // Data 구조체는 main 패키지에 속해있기 때문에 main.Data 로 출력됨.

	fNumType := reflect.TypeOf(fNum)   // fNum의 타입 정보
	fNumValue := reflect.ValueOf(fNum) // fNum의 값 정보
	fmt.Println("**** fNumType 정보")
	fmt.Println("fNumType ->", fNumType, "/ fNumValue ->", fNumValue)
	fmt.Println("type.Name() ->", fNumType.Name(), "/ type.Size() ->", fNumType.Size(), "/ type.Kind() ->", fNumType.Kind())
	fmt.Println("fNumType == reflect.Float64 -> ", fNumType.Kind() == reflect.Float64, "/ fNumType == reflect.Float32 -> ", fNumType.Kind() == reflect.Float32)

	fmt.Println("**** fNumValue 정보")
	fmt.Println("fNumValue == reflect.Float64 -> ", fNumValue.Kind() == reflect.Float64, "/ fNumValue == reflect.Float32 -> ", fNumValue.Kind() == reflect.Float32)
	fmt.Println("각 타입에 맞는 함수를 사용하면 값을 가져올 수 있다. value.Float() -> ", fNumValue.Float())
	// fNumValue.Int() // panic이 발생한다. panic: reflect: call of reflect.Value.Int on float32 Value
}

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func example_1_structTag() {
	p := Person{}

	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))
}

func example_2_pointerAndInterface() {
	fmt.Println("**** pointer variable")
	var a *int = new(int)
	*a = 1

	fmt.Println("a 변수의 타입 ->", reflect.TypeOf(a), "/ a 변수의 값. 포인터이기 때문에 주소값이다. ->", reflect.ValueOf(a))
	// fmt.Println(reflect.ValueOf(a).Int()) // panic 발생. reflect: call of reflect.Value.Int on ptr Value

	fmt.Println(reflect.ValueOf(a).Elem())       // 포인터의 메모리에 저장된 실제 값 정보를 가져온다.
	fmt.Println(reflect.ValueOf(a).Elem().Int()) // 값을 Int 형으로 가져옴.

	fmt.Println("**** interface{} variable")
	var b interface{}
	b = 1

	fmt.Println("b 변수의 타입 ->", reflect.TypeOf(b), "/ b변수의 값. ->", reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(b).Int()) // 값이 나온다.

	// fmt.Println(reflect.ValueOf(b).Elem()) // panic 발생. reflect: call of reflect.Value.Elem on int Value
	//fmt.Println(reflect.ValueOf(b).Elem().Int()) // panic 발생. reflect: call of reflect.Value.Elem on int Value

}
