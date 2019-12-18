package main

import (
  "fmt"
  "os"
)

func main() {
  argsWithProgramCmd := os.Args
  argsWithoutProgramCmd := os.Args[1:]

  fmt.Println("argsWithProgramCmd ->", argsWithProgramCmd)
  fmt.Println("argsWithoutProgramCmd  ->", argsWithoutProgramCmd)
}

