package main

import "fmt"

func main() {
	// ! Shadow variables
	// x := 10
	// if x > 5 {
	// 	fmt.Println(x)
	// 	x = 5
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// x := 10
	// if x > 5 {
	// 	x, y := 5, 20
	// 	fmt.Println(x, y)
	// }
	// fmt.Println(x)

	// ? Shadow package import
	// x := 10
	// fmt.Println(x)
	// fmt := "oops"
	// fmt.Println(fmt)

	// ? Detecting Shadowed Variables
	// fmt.Println(false)
	// make := 10
	// fmt.Println(make)

	// ! If
	// n := rand.Intn(10)
	// if n == 0 {
	// 	fmt.Println("That's too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big:", n)
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }

	// if n := rand.Intn(10); n == 0 {
	// 	fmt.Println("That's too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big:", n)
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }

	// ! FOR, FOUR WAYS
	// ? A complete C-style for
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// ? A condition-only for
	// i := 2
	// for i < 20 {
	// 	fmt.Println(i)
	// 	i *= 2
	// }
	// ? An infinite for
	// for {
	// 	fmt.Println("Hello")
	// }

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		if i%5 == 0 {
	// 			fmt.Println("FizzBuzz")
	// 		} else {
	// 			fmt.Println("Fizz")
	// 		}
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }
	// ? for-range
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	fmt.Println(v)
	// }

	// uniqueNames := map[string]bool{
	// 	"Fred": true, "Raul": true, "Wilma": true,
	// }

	// for k := range uniqueNames {
	// 	fmt.Println(k)
	// }

	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Loop", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	// samples := []string{"hello", "apple_π!"}
	// for _, sample := range samples {
	// 	for i, r := range sample {
	// 		fmt.Println(i, r, string(r))
	// 	}
	// 	fmt.Println()
	// }
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }
	// fmt.Println(evenVals)

	// ? Labelling loop
	// 	samples := []string{"hello", "apple_π!"}
	// outer:
	// 	for _, sample := range samples {
	// 		for i, r := range sample {
	// 			fmt.Println(i, r, string(r))
	// 			if r == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}
	// name := "Chibao"
	// for i, r := range name {
	// 	fmt.Println(i, r, string(r))
	// }

	// ! Switch
	// words := []string{"a", "cow", "smile", "gopher",
	// 	"octopus", "anthropologist"}
	// for _, word := range words {
	// 	switch size := len(word); size {
	// 	case 1, 2, 3, 4:
	// 		fmt.Println(word, "is a short word!")
	// 	case 5:
	// 		wordLen := len(word)
	// 		fmt.Println(word, "is exactly the right length:", wordLen)
	// 	case 6, 7, 8, 9:
	// 	default:
	// 		fmt.Println(word, "is a long word!")
	// 	}
	// }
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		wordLen := len(word)
		switch {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}
