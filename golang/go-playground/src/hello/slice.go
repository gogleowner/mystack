package example

import . "fmt"

func main() {
	var a []int = make([]int, 5, 10)
	// a := make([]int, 5, 10)

	Println("length :", len(a), "capacity :", cap(a))

	b := []int{1, 2, 3}
	b = append(b, 4, 5, 6)

	Println(b)

	c := make([]int, 3)
	copy(c, b)
	Println("c array size is 3 ->", c, " / b array range is 1~6", b)

	// partition slice
	d := b[1:6]
	Println(d)

	e := b[0:4:len(b)]
	Println(e, "length:", len(e), "capacity:", cap(e))
}
