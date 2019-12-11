package main

import (
	"fmt"
	"strings"
)

func main() {
	example_1_searchString()
}

func example_1_searchString() {
	fmt.Println(strings.Contains("hello world", "wo"))
	fmt.Println(strings.Contains("hello world", "w o"))
	fmt.Println(strings.Contains("hello world", "ow"))
	fmt.Println(strings.ContainsAny("hello world", "wo"))
	fmt.Println(strings.ContainsAny("hello world", "w o"))
	fmt.Println(strings.ContainsAny("hello world", "ow"))
	fmt.Println(strings.Count("Hello Helium", "He"))

	var r rune = '하'
	fmt.Println(strings.ContainsRune("안녕하세요", r))
	fmt.Println(strings.HasPrefix("Hello world!", "He"))
	fmt.Println(strings.HasSuffix("Hello world!", "rld!"))
}

