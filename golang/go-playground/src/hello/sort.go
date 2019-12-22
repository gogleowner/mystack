package main

import (
	"fmt"
	"sort"
)

func main() {
	example_1_sort()
}

func example_1_sort() {
	a := []int{10, 5, 3, 7, 6 }
	sort.Sort(sort.IntSlice(a)) // 오름차순 정렬
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a))) // 내림차순 정렬
	fmt.Println(a)
	sort.Ints(a) // 오름차순 정렬
	fmt.Println(a)

	b := []float64{4.2, 7.6, 5.5, 1.3, 9.9}
	sort.Sort(sort.Float64Slice(b))
	fmt.Println(b)
	sort.Sort(sort.Reverse(sort.Float64Slice(b)))
	fmt.Println(b)
	sort.Float64s(b)
	fmt.Println(b)

	c := []string{"Maria", "Andrew", "John"}
	sort.Sort(sort.StringSlice(c))
	fmt.Println(c)
	sort.Sort(sort.Reverse(sort.StringSlice(c)))
	fmt.Println(c)
	sort.Strings(c)
	fmt.Println(c)
}

