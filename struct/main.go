package main

import "fmt"

func main() {

	// ! Struct Equality
	// a1 := Author{
	// 	name:      "Moana",
	// 	branch:    "CSE",
	// 	language:  "Python",
	// 	Particles: 38,
	// }

	// a2 := Author{
	// 	name:      "Moana",
	// 	branch:    "CSE",
	// 	language:  "Python",
	// 	Particles: 38,
	// }

	// a3 := Author{
	// 	name:      "Dona",
	// 	branch:    "CSE",
	// 	language:  "Python",
	// 	Particles: 38,
	// }
	// if a1 == a2 {

	// 	fmt.Println("Variable a1 is equal to variable a2")

	// } else {

	// 	fmt.Println("Variable a1 is not equal to variable a2")
	// }
	// if a2 == a3 {

	// 	fmt.Println("Variable a2 is equal to variable a3")

	// } else {

	// 	fmt.Println("Variable a2 is not equal to variable a3")
	// }
	// fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(a1, a2))

	// // Comparing a2 with a3
	// // Using DeepEqual() method
	// fmt.Println("Is a2 equal to a3: ", reflect.DeepEqual(a2, a3))
	// p := Person{
	// 	name:    "Ly Minh Thang",
	// 	age:     24,
	// 	address: Address{street: "Bac Lieu", city: "Bac Lieu"},
	// }
	// fmt.Println(p)

	// ! Anonymous Structure

	// person := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "Chi bao",
	// 	age:  14,
	// }
	// fmt.Println(person)

	// student := Student{
	// 	personalDetails: struct {
	// 		name       string
	// 		enrollment int
	// 	}{
	// 		name:       "Bao",
	// 		enrollment: 2,
	// 	},
	// 	GPA: 3.8,
	// }
	// fmt.Println(student)

	// student := Student1{12345, "A", 3.8}
	// fmt.Println("Enrollment:", student.int)
	// fmt.Println("Name:", student.string)
	// fmt.Println("GPA:", student.float64)

	// ! Promoted methods
	// values := employee{
	// 	post: "HR",
	// 	eid:  4567,
	// 	details: details{

	// 		name:    "Sumit",
	// 		age:     28,
	// 		gender:  "Male",
	// 		psalary: 890,
	// 	},
	// }
	// // Promoted fields of the
	// // employee structure
	// fmt.Println("Name: ", values.name)
	// fmt.Println("Age: ", values.age)
	// fmt.Println("Gender: ", values.gender)
	// fmt.Println("Per day salary: ", values.psalary)

	// // Promoted method of the
	// // employee structure
	// fmt.Println("Total Salary: ", values.promotedMethod(30))

	// // Normal fields of
	// // the employee structure
	// fmt.Println("Post: ", values.post)
	// fmt.Println("Employee id: ", values.eid)

	// Initializing the fields of
	// the employee structure
	// values := employee{
	// 	post:     "HR",
	// 	eid:      4567,
	// 	particle: 5,
	// 	details: details{

	// 		name:    "Sumit",
	// 		age:     28,
	// 		gender:  "Male",
	// 		psalary: 890,
	// 	},
	// }

	// // Promoted fields of
	// // the employee structure
	// fmt.Println("Name: ", values.name)
	// fmt.Println("Age: ", values.age)
	// fmt.Println("Gender: ", values.gender)
	// fmt.Println("Per day salary: ", values.psalary)

	// // Promoted method of
	// // the employee structure
	// fmt.Println("Total Salary: ", values.promotmethod(30))

	// // Normal fields of
	// // the employee structure
	// fmt.Println("Post: ", values.post)
	// fmt.Println("Employee id: ", values.eid)
	// fmt.Println("Total Articles: ", values.promotmethod(30))

	var p = Person{
		Name: "Cho bao",
	}
	p.Greet = func() string {
		return "Hello " + p.Name
	}
	fmt.Println(p.Greet())
}

// ! Function as field in Struct
type Person struct {
	Name  string
	Greet func() string
}

// ! Promoted Methods

// Structure
type details struct {

	// Fields of the
	// details structure
	name    string
	age     int
	gender  string
	psalary int
}

// Method 1
func (e employee) promotmethod(tarticle int) int {
	return e.particle * tarticle
}

// Nested structure
type employee struct {
	post     string
	particle int
	eid      int
	details
}

// Method 2
func (d details) promotmethod(tsalary int) int {
	return d.psalary * tsalary
}

// type details struct {

// 	// Fields of the
// 	// details structure
// 	name    string
// 	age     int
// 	gender  string
// 	psalary int
// }

// // Nested structure
// type employee struct {
// 	post string
// 	eid  int
// 	details
// }

func (d details) promotedMethod(tsalary int) int {
	return d.psalary * tsalary
}

// ! Anonymous field
type Student1 struct {
	int
	string
	float64
}

// ! Anonymous struct
type Student struct {
	personalDetails struct { // Anonymous inner structure for personal details
		name       string
		enrollment int
	}
	GPA float64 // Standard field
}

// type Person struct {
// 	name    string
// 	age     int
// 	address Address
// }

// In the above, the type keyword introduces a new type
type Address struct {
	street string
	city   string
}

type Author struct {
	name      string
	branch    string
	language  string
	Particles int
}

// ? More compact

// type Address struct {
//     name, street, city, state string
//     Pincode int
// }
