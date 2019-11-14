package main

import . "fmt"

func main() {
  // definition
  var a = [5]int{32,29,78,16,81}

  Println(a)

  var b = [...]int {
    1,
    3,
    5,
    6,
    9,
  }

  Println(b)

  // iterate
  for i := 0; i < len(a); i++ {
    Println(a[i])
  }

  for idx, value := range a {
    Println(idx, value)
  }

  for idx := range a {
    Print(idx, " ")
  }
  Println()
  for _, value := range a {
    Print(value, " ")
  }
}
