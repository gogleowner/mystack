package main

import (
  "fmt"
  "reflect"
)

func h(args []reflect.Value) []reflect.Value {
  fmt.Println("Hello world!")
  return nil
}

func main() {
  var hello func()

  fn := reflect.ValueOf(&hello).Elem() // hello의 주소를 넘기고, Elem으로 값 정보를 가져온다.

  v := reflect.MakeFunc(fn.Type(), h) // h의 함수 정보 생성

  fmt.Println("Before fn.Set()", fn)
  fn.Set(v) // hello 의 값 정보인 fn에 h의 함수정보 v를 설정하여 함수를 연결

  fmt.Println("After fn.Set()", fn)
  hello()
}

// 나머지 예제는 나중에..

