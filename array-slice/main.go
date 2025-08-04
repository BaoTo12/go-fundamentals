package main

import (
	"bytes"
	"fmt"
)

// ? How to create an array in Golang
/// Var array_name[length]Type
/// Var array_name[index] = element
/// array_name:= [length]Type{item1, item2, item3,...itemN}
/// Array_name[Length1][Length2]..[LengthN]Type

func main() {
	// scores := [6]int{67, 59, 29, 35, 4, 34}
	// average := calculateAverage(scores, len(scores))
	// fmt.Printf("%d\n", average)

	// values := [5]int{1, 2, 3, 4, 5}

	// // Modifying the array
	// incrementArray(values)

	// fmt.Println("Incremented array:", values)

	// slice := []int{1, 2, 3, 4, 5, 6}
	// newSclice := append(slice, 4, 5, 6)
	// fmt.Println(slice)
	// fmt.Println(newSclice)

	// Creating an array
	// arr := [7]string{"This", "is", "the", "tutorial",
	// 	"of", "Go", "language"}

	// // Display array
	// fmt.Println("Array:", arr)

	// // Creating a slice
	// myslice := arr[1:6]

	// // Display slice
	// fmt.Println("Slice:", myslice)

	// // Display length of the slice
	// fmt.Printf("Length of the slice: %d", len(myslice))

	// // Display the capacity of the slice
	// fmt.Printf("\nCapacity of the slice: %d", cap(myslice))

	// arr := [7]string{"This", "is", "the", "tutorial",
	// 	"of", "Go", "language"}

	// slice := arr[:]
	// fmt.Println(slice)

	// Creating s slice
	// oRignAl_slice := []int{90, 60, 40, 50, 34, 49, 30}

	// // Creating slices from the given slice
	// var my_slice_1 = oRignAl_slice[1:5]
	// my_slice_2 := oRignAl_slice[0:]
	// my_slice_3 := oRignAl_slice[:6]
	// my_slice_4 := oRignAl_slice[:]
	// my_slice_5 := my_slice_3[2:4]

	// // Display the result
	// fmt.Println("Original Slice:", oRignAl_slice)
	// fmt.Println("New Slice 1:", my_slice_1)
	// fmt.Println("New Slice 2:", my_slice_2)
	// fmt.Println("New Slice 3:", my_slice_3)
	// fmt.Println("New Slice 4:", my_slice_4)
	// fmt.Println("New Slice 5:", my_slice_5)

	// slice := make([]int, 4, 5)
	// fmt.Println(slice)

	// var my_slice_2 = make([]int, 7)
	// fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
	// 	my_slice_2, len(my_slice_2), cap(my_slice_2))

	// ! nil slice
	// var myslice []int
	// fmt.Println(myslice)
	// fmt.Printf("Length = %d\n", len(myslice))
	// fmt.Printf("Capacity = %d ", cap(myslice))

	// ! modifying an slice
	// arr := [6]int{55, 66, 77, 88, 99, 22}
	// slc := arr[0:4]

	// Before modifying

	// fmt.Println("Original_Array: ", arr)
	// fmt.Println("Original_Slice: ", slc)

	// // After modification
	// slc[0] = 100
	// slc[1] = 1000
	// slc[2] = 1000

	// fmt.Println("\nNew_Array: ", arr)
	// fmt.Println("New_Slice: ", slc)

	//! slice comparison
	// s1 := []int{12, 34, 56}
	// var s2 []int

	// fmt.Println(s1 == nil)
	// fmt.Println(s2 == nil)

	// ! copy slice
	// var source = []int{10, 20, 30, 40, 50}

	// destination := make([]int, len(source))

	// count := copy(destination, source)
	// fmt.Println("Source:", source)
	// fmt.Println("Destination:", destination)
	// fmt.Println("Elements copied:", count)

	// destination := []int{}
	// destination = append(destination, source...)

	// fmt.Println("Source:", source)
	// fmt.Println("Destination:", destination)

	// ! modify a slice
	// number := []int{1, 2, 3, 4, 5, 6}
	// printSlice(number)
	// modifySlice(number)
	// printSlice(number)

	// ! return a slice from a function
	// a := []int{1, 2, 3}

	// newNumbers := addElement(a, 4)

	// fmt.Println("Original slice:", a)
	// fmt.Println("New slice:", newNumbers)

	// a := []int{1, 2, 3}
	// fmt.Println("Original slice:", a)

	// // Passing slice pointer to the function
	// modifySlicePointer(&a)

	// fmt.Println("Modified slice:", a)

	// ! Compare slices of bytes
	// slice1 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	// slice2 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	// slice3 := []byte{'g', 'o', 'L', 'A', 'N', 'g'}
	// fmt.Println(bytes.Compare(slice1, slice2)) // Output: 0 (equal)
	// fmt.Println(bytes.Compare(slice1, slice3)) // Output: -1 (slice1 < slice3)

	// slc1 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	// slc2 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	// slc3 := []byte{'g', 'o', 'L', 'A', 'N', 'g'}

	// fmt.Println(bytes.Equal(slc1, slc2)) // Output: true
	// fmt.Println(bytes.Equal(slc1, slc3)) // Output: false

	// slc1 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	// slc2 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	// slc3 := []byte{'g', 'o', 'L', 'A', 'N', 'g'}

	// fmt.Println(reflect.DeepEqual(slc1, slc2)) // Output: true
	// fmt.Println(reflect.DeepEqual(slc1, slc3)) // Output: false

	// ! sorting slice
	// intSlice := []int{42, 23, 16, 15, 8, 4}
	// fmt.Println("Before:", intSlice)
	// sort.Ints(intSlice)
	// fmt.Println("After:", intSlice)
	// fmt.Println("sorted after sorting?:", sort.IntsAreSorted(intSlice))

	// intSlice := []int{42, 23, 16, 15, 8, 4}
	// sort.Slice(intSlice, func(i, j int) bool {
	// 	return intSlice[i] > intSlice[j]
	// })
	// fmt.Println("After:", intSlice)

	// ! Trim slice
	// Creating and initializing
	// the slice of bytes
	// Using shorthand declaration
	// slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f',
	// 	'o', 'r', 'G', 'e', 'e', '#', 's', '#', '#'}

	// slice_2 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}

	// slice_3 := []byte{'%', 'g', 'e', 'e', 'k', 's', '%'}

	// // Displaying slices
	// fmt.Println("Original Slice:")
	// fmt.Printf("Slice 1: %s", slice_1)
	// fmt.Printf("\nSlice 2: %s", slice_2)
	// fmt.Printf("\nSlice 3: %s", slice_3)

	// // Trimming specified leading
	// // and trailing Unicodes points
	// // from the given slice of bytes
	// // Using Trim function
	// res1 := bytes.Trim(slice_1, "!#")
	// res2 := bytes.Trim(slice_2, "*^")
	// res3 := bytes.Trim(slice_3, "@")

	// // Display the results
	// fmt.Printf("New Slice:\n")
	// fmt.Printf("\nSlice 1: %s", res1)
	// fmt.Printf("\nSlice 2: %s", res2)
	// fmt.Printf("\nSlice 3: %s", res3)

	// res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
	// res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
	// res3 := bytes.Trim([]byte("^^Geek&&"), "$")

	// // Display the results
	// fmt.Printf("Final Slice:\n")
	// fmt.Printf("\nSlice 1: %s", res1)
	// fmt.Printf("\nSlice 2: %s", res2)
	// fmt.Printf("\nSlice 3: %s", res3)

	// ! Split slice of bytes
	// Initial byte slice
	// data := []byte("a,b,c")
	// fmt.Println(data)
	// // Splitting the byte slice
	// parts := bytes.Split(data, []byte(","))
	// fmt.Println("Split parts:")
	// for _, part := range parts {
	// 	fmt.Println(string(part))
	// }

	data := []byte("a,b,c")

	// Attempt to split with a non-existent separator
	parts := bytes.Split(data, []byte("|"))
	fmt.Println("Split parts with non-existent separator '|':")
	for _, part := range parts {
		fmt.Println(string(part))
	}
}

func modifySlicePointer(slice *[]int) {
	*slice = append(*slice, 6)
}

func addElement(slice []int, element int) []int {
	return append(slice, element)
}

func modifySlice(slice []int) {
	for i := range slice {
		slice[i] *= 2 // Doubling each element
	}
}
func printSlice(slice []int) {
	for _, val := range slice {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func calculateAverage(arr [6]int, size int) int {
	var total int
	for _, val := range arr {
		total += val
	}
	return total / size
}
func incrementArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i]++
	}
}
