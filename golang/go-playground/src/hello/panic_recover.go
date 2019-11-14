package main

import . "fmt"

func main() {
  f()
  Println("Hello, world")
}

func f() {
  defer func() {
    s := recover()

    Println(s)
  }()

  panic("Error!")
}

