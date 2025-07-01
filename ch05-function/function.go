package main

import (
	"errors"
	"fmt"
)

func main() {
	// result := div(5, 2)
	// fmt.Println(result)

	// fmt.Println(addTo(3))
	// fmt.Println(addTo(3, 2))
	// fmt.Println(addTo(3, 2, 4, 6, 8))
	// a := []int{4, 3}
	// fmt.Println(addTo(3, a...))
	// fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	// result, remainder, _ := divAndRemainder(5, 2)

	// fmt.Println(result, remainder)

	// ! Functions Are Values

	// opMap := map[string]func(int, int) int{
	// 	"+": add,
	// 	"-": sub,
	// 	"*": mul,
	// 	"/": div,
	// }

	// expressions := [][]string{
	// 	[]string{"2", "+", "3"},
	// 	[]string{"2", "-", "3"},
	// 	[]string{"2", "*", "3"},
	// 	[]string{"2", "/", "3"},
	// 	[]string{"2", "%", "3"},
	// 	[]string{"two", "+", "three"},
	// 	[]string{"5"},
	// }

	// for _, expression := range expressions {
	// 	if len(expression) != 3 {
	// 		fmt.Println("invalid expression:", expression)
	// 		continue
	// 	}
	// 	p1, err := strconv.Atoi(expression[0])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	op := expression[1]
	// 	opFunc, ok := opMap[op]
	// 	if !ok {
	// 		fmt.Println("unsupported operator:", op)
	// 		continue
	// 	}
	// 	p2, err := strconv.Atoi(expression[2])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	result := opFunc(p1, p2)
	// 	fmt.Println(result)
	// }

	// ! Function Type Declarations
	// type opFuncType func(int, int) int
	// var opMap = map[string]opFuncType {

	// }

	// ! Anonymous function
	// for i := range 5 {
	// 	func(j int) {
	// 		fmt.Println("printing", j, "from inside of an anonymous function")
	// 	}(i)
	// }

	// ! Passing Functions as Parameters
	// type Person struct {
	// 	FirstName string
	// 	LastName  string
	// 	Age       int
	// }
	// people := []Person{
	// 	{"Pat", "Patterson", 37},
	// 	{"Tracy", "Bobbert", 23},
	// 	{"Fred", "Fredson", 18},
	// }
	// fmt.Println(people)

	// sort.Slice(people, func(i int, j int) bool {
	// 	return people[i].LastName < people[j].LastName
	// })
	// fmt.Println(people)
	// sort.Slice(people, func(i int, j int) bool {
	// 	return people[i].Age < people[j].Age
	// })
	// fmt.Println(people)

	// ! Returning Functions from Functions
	// twoBase := makeMul(2)
	// threeBase := makeMul(3)

	// for v := range 3 {
	// 	fmt.Println(twoBase(v), threeBase(v))
	// }

	// ! Defer
	// if len(os.Args) < 2 {
	// 	log.Fatal("no file specified")
	// }
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// data := make([]byte, 2048)
	// for {
	// 	count, err := f.Read(data)
	// 	os.Stdout.Write(data[:count])
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	// }

	// demonstrateDefer()

	// ! Go Is Call By Value
	// p := person{}
	// i := 2
	// s := "Hello"
	// modifyFails(i, s, p)
	// fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)
	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)

}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
	delete(m, 2)
	m[4] = "ChiBao"
}
func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}

	s = append(s, 10)
}

// type person struct {
// 	age  int
// 	name string
// }

// func modifyFails(i int, s string, p person) {
// 	i = i * 2
// 	s = "Goodbye"
// 	p.name = "Bob"
// }

func demonstrateDefer() {
	fmt.Println("Function starts")

	defer fmt.Println("First defer - will run LAST")
	defer fmt.Println("Second defer - will run second")
	defer fmt.Println("Third defer - will run FIRST")

	fmt.Println("Function middle")
	fmt.Println("Function about to return")
	return
	// Deferred functions run here, in reverse order
}

func makeMul(base int) func(int) int {
	return func(factor int) int {
		return factor * base
	}
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

func divAndRemainder(numerator, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// func div(numerator int, denominator int) int {
// 	return numerator / denominator
// }
