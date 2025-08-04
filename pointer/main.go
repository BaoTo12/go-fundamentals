package main

import "fmt"

func main() {
	// var name *dataType
	// var a = 45
	// var pointer = &a
	// pointer2 := &a
	// fmt.Println(pointer)
	// fmt.Println(pointer2)
	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// updateSlice(arr[:])
	// fmt.Println(arr)
	// slice := make([]int, 3)
	// fmt.Println(slice)

	// ! Pointer to struct
	// p := Person{
	// 	name: "Chi Bao",
	// 	age:  12,
	// }
	// fmt.Println(p)
	// updateStruct(&p)
	// fmt.Println(p)

	// p1 := Person{name: "A", age: 25}

	// // Creating a pointer to the struct
	// personPointer := &p1

	// fmt.Println("Name:", personPointer.name) // Automatically dereferences
	// fmt.Println("Age:", personPointer.age)

	// // Modifying struct values using the pointer
	// personPointer.age = 26
	// fmt.Println("Updated struct:", p1)

	// personPointer := new(Person)
	// personPointer.name = "Thang"
	// personPointer.age = 25
	// fmt.Println(personPointer)

	// ! Comparing pointers
	// var pointer1 *int
	// var pointer2 *int
	// fmt.Println(pointer1 == pointer2)

	val1 := 2345
	val2 := 567

	// Creating and initializing pointers
	var p1 *int
	p1 = &val1
	p2 := &val2
	p3 := &val1

	res1 := &p1 == &p2
	fmt.Println("Is p1 pointer is equal to p2 pointer: ", res1)

	res2 := p1 == p2
	fmt.Println("Is p1 pointer is equal to p2 pointer: ", res2)

	res3 := p1 == p3
	fmt.Println("Is p1 pointer is equal to p3 pointer: ", res3)

	res4 := p2 == p3
	fmt.Println("Is p2 pointer is equal to p3 pointer: ", res4)

	res5 := &p3 == &p1
	fmt.Println("Is p3 pointer is equal to p1 pointer: ", res5)

}
func updateStruct(p *Person) {
	p.name = "Thang"
}

type Person struct {
	name string
	age  int
}

func updateArray(arr *[5]int, index int) {
	(*arr)[index] = 4
}
func updateSlice(slice []int) {
	slice[2] = 700
}
