package main

import . "fmt"

func main() {
	aMap := make(map[string]int)
	aMap["a"] = 1
	aMap["b"] = 1

	Println(aMap)

	bMap := map[string]int{
		"hello": 1,
		"world": 2,
	}
	Println(bMap)
	Println(bMap["hello"])

	Println(bMap["blabla"]) // 0

	// contains
	value, isExist := bMap["blabla"]
	Println(value, isExist) // 0, false

	if value, isExist := bMap["blabla"]; isExist {
		Println("isExist!!", value)
	} else {
		Println("not exist!!")
	}

	// iterate
	for key, value := range bMap {
		Println(key, value)
	}

	for _, value := range bMap {
		println(value)
	}

	for key := range bMap {
		println(key)
	}

	// delete
	delete(bMap, "world")
	Println(bMap)

	// map in map
	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022e+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676e+24,
			"orbitalPeriod": 224.70069,
		},
	}
	Println(terrestrialPlanet)
	Println(terrestrialPlanet["Venus"]["meanRadius"])
}
