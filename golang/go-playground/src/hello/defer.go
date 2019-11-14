package main

import . "fmt"

func main() {
  // defer : 특정 함수를 현재 함수가 끝나기 직전에 실행하는 기능
  defer hello("last")

  hello("first")

  HelloWorld2()
}

func hello(name string) {
  Println("hello", name)
}

func HelloWorld() {
  defer func() {
    Print(" world")
  }()
  func() {
    Print("Hello")
  }()
}

func HelloWorld2() {
  defer func() {
    Print("a")
  }()
  defer func() {
    Print("b")
  }()
  defer func() {
    Print("c")
  }()
  func() {
    Print("Hello")
  }()
}

