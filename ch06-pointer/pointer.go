package main

import "fmt"

func main() {
	// var x int32 = 10
	// var y bool = true

	// pointerX := &x
	// pointerY := &y

	// var str *string

	// fmt.Println(pointerX, pointerY)
	// fmt.Println(*pointerX, *pointerY)
	// fmt.Println(str)

	// ! pointer type
	// x := 10
	// var pointerX *int
	// pointerX = &x
	// fmt.Println(*pointerX)

	s := make([]int, 3, 10)
	fmt.Println(s)
	modSlice(s)
	fmt.Println(s)
}

func modMap(m map[string]int) {
	m["Hello"] = 30
}

func modSlice(s []int) {
	s[2] = 4
	s = append(s, 10)
	fmt.Println(s)
}
