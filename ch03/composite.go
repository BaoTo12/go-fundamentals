package main

import "fmt"

func main() {
	// var x = [...]int{1, 2, 3}
	// var y = [3]int{1, 2, 3}

	// fmt.Println(x == y)
	// var x = [...]int{1, 2, 3}

	// x[0] = 10
	// fmt.Println(x)
	// fmt.Println(len(x))

	// var x = []int{1, 5: 4, 10: 12}
	// fmt.Println(x)

	// var x = []int{1, 4, 11}
	// var y = []int{7, 10}
	// x = append(x, y...)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// fmt.Println(x, len(x), cap(x))

	// x := make([]int, 5, 10)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// x := make([]int, 0, 10)
	// x = append(x, 5, 6, 7, 8)
	// fmt.Println(x)

	// x := []int{1, 2, 3, 4, 5, 6, 7}
	// y := x[:2]
	// z := x[2:]
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// y = append(y, 99)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	// x := []int{1, 2, 3, 4}
	// y := x[:2]
	// fmt.Println(cap(x), cap(y))
	// y = append(y, 30)
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	// x := make([]int, 0, 10)
	// y := x[3:9]
	// fmt.Println(cap(x), cap(y))

	// !
	// x := make([]int, 0, 5)
	// x = append(x, 1, 2, 3, 4)
	// y := x[:2]
	// z := x[2:]
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, 30, 40, 50)
	// x = append(x, 60)
	// z = append(z, 70)
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// x := make([]int, 0, 5)
	// x = append(x, 1, 2, 3, 4)
	// y := x[:2:2]
	// z := x[2:4:4]
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, 30)
	// x = append(x, 60)
	// z = append(z, 70)
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// x := [4]int{5, 6, 7, 8}
	// y := x[:2]
	// z := x[2:]
	// x[0] = 10
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// ! COPY
	// x := []int{1, 2, 3, 4}
	// y := make([]int, 2)
	// copy(y, x[2:])
	// fmt.Println(y)

	// x := []int{1, 2, 3, 4}
	// copy(x[:3], x[1:])
	// fmt.Println(x)
	// src: 2, 3, 4
	// des: 1, 2, 3

	// x := []int{1, 2, 3, 4}
	// d := [4]int{5, 6, 7, 8}
	// y := make([]int, 2)
	// copy(y, d[:])
	// fmt.Println(y)
	// copy(d[:], x)
	// fmt.Println(d)

	//! String
	// var s string = "Hello there"
	// var b byte = s[6]
	// fmt.Println(b)
	// var s string = "Hello "
	// fmt.Println(len(s))

	// var a rune = 'a'
	// var str1 string = string(a)
	// fmt.Println(str1)
	// var b byte = 'b'
	// var str2 string = string(b)

	// fmt.Println(str2)

	// var a int = 65
	// var y = string(a)
	// fmt.Println(y)

	// var s string = "Hello Word"
	// var bs []byte = []byte(s)
	// fmt.Println(bs)
	// var rs []rune = []rune(s)
	// fmt.Println(rs)

	// ! map
	// var nilMap map[string]int
	// fmt.Println(nilMap["a"])

	// totalWins := map[string]int{}

	// teams := map[string][]string{
	// 	"Orcas": []string{"Fred", "Peter"},
	// 	"Lions": []string{"Waldo", "Raul"},
	// }

	// fmt.Println(teams)

	// ages := make(map[int][]string , 10)
	// totalWins := map[string]int{}
	// totalWins["Orcas"] = 1
	// totalWins["Lions"] = 2
	// fmt.Println(totalWins["Orcas"])
	// fmt.Println(totalWins["Kittens"])
	// fmt.Println(totalWins)
	// totalWins["Kittens"]++
	// fmt.Println(totalWins["Kittens"])
	// totalWins["Lions"] = 3
	// fmt.Println(totalWins["Lions"])
	// fmt.Println(totalWins)

	// m := map[string]int{
	// 	"hello": 5,
	// 	"world": 0,
	// }
	// v, ok := m["hello"]
	// fmt.Println(v, ok)
	// v, ok = m["world"]
	// fmt.Println(v, ok)
	// v, ok = m["goodbye"]
	// fmt.Println(v, ok)

	// ! Set
	// intSet := map[int]bool{}
	// vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	// for _, v := range vals {
	// 	fmt.Print(v, " ")
	// 	intSet[v] = true
	// }
	// fmt.Println(len(vals), len(intSet))
	// fmt.Println(intSet[5])
	// fmt.Println(intSet[500])
	// if intSet[100] {
	// 	fmt.Println("100 is in the set")
	// }
	// fmt.Println(intSet)

	// ! struct
	// type person struct {
	// 	name string
	// 	age  int
	// }

	// var personA person
	// fmt.Println(personA)
	// // ? struct literal
	// bob := person{
	// 	"Julia",
	// 	40,
	// }

	// fmt.Println(bob)
	// beth := person{
	// 	age:  10,
	// 	name: "Hola",
	// }
	// fmt.Println(beth)

	// ? anonymous struct
	// var person struct {
	// 	name string
	// 	age int
	// }

	// beth := struct {
	// 	age  int
	// 	name string
	// }{
	// 	age:  10,
	// 	name: "Hola",
	// }
	// fmt.Println(beth)

	// ben := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "Ben",
	// 	age:  24,
	// }
	// fmt.Println(ben)

	// ? Comparing an convert structs
	// type firstPerson struct {
	// 	name string
	// 	age  int
	// }
	// type secondPerson struct {
	// 	name string
	// 	age  int
	// }
	// type thirdPerson struct {
	// 	age  int
	// 	name string
	// }
	// type fourthPerson struct {
	// 	firstName string
	// 	age       int
	// }
	// type fifthPerson struct {
	// 	name          string
	// 	age           int
	// 	favoriteColor string
	// }

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
}
