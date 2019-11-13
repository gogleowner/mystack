package main

import . "fmt"

func main() {
	for i := 99; i >= 0; i-- {
		switch {
		case i >= 2:
			Printf("%d bottles of beer on the wall, %d bottles of beer\n", i, i)
			bottleStr := "bottles"
			if i-1 == 1 {
				bottleStr = "bottle"
			}
			Printf("Take one down, pass it aroud, %d  %s of beer on the wall\n", i-1, bottleStr)
		case i == 1:
			Println("1 bottle of beer on the wall, 1 bottle of beer.")
			Println("Take one down, pass it aroud, No more bottles of beer on the wall.")
		default:
			Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
		}
	}
}
