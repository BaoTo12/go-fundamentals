package main

import "fmt"

func main() {
	// var s Shape
	// s = Circle{radius: 5}
	// fmt.Println("C Area:", s.Area())
	// fmt.Println("C Perimeter:", s.Perimeter())

	// s = Rectangle{length: 4, width: 3}
	// fmt.Println("R Area:", s.Area())
	// fmt.Println("R Perimeter:", s.Perimeter())

	// ! Type assertion
	// s := Square{
	// 	side: 2,
	// }
	// printArea(s)

	// ! Multiple Interfaces
	// values := author{
	// 	a_name:    "Mickey",
	// 	branch:    "Computer science",
	// 	college:   "XYZ",
	// 	year:      2012,
	// 	salary:    50000,
	// 	particles: 209,
	// 	tarticles: 309,
	// }

	// // Accessing the method
	// // of the interface 1
	// var i1 AuthorDetails = values
	// i1.details()

	// // Accessing the method
	// // of the interface 2
	// var i2 AuthorArticles = values
	// i2.articles()

	// doc := &Document{content: "Initial content"}

	// // using the Reader interface to read the document content
	// var r Reader = doc
	// fmt.Println("Content before writing:", r.Read())

	// // using the Writer interface to write new content to the document
	// var w Writer = doc
	// w.Write("New content")

	// // using the ReadWriter interface to read the updated document content
	// var rw ReadWriter = doc
	// fmt.Println("Content after writing:", rw.Read())
	// fmt.Println("Content before writing:", r.Read())

	// ! Embedding Interfaces
	// to the structure
	values := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}
	s := values
	s.articles()
	s.details()
}

// Implementing method of
// the interface 1
func (a author) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d",
		a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}

// Implementing method
// of the interface 2
func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

// Interface 1
type AuthorDetails interface {
	details()
}

// Interface 2
type AuthorArticles interface {
	articles()
}

// Interface 3

// Interface 3 embedded with
// interface 1 and 2
type FinalDetails interface {
	AuthorDetails
	AuthorArticles
}

// Structure
type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

// type Reader interface {
// 	Read() string
// }

// type Writer interface {
// 	Write(string)
// }

// type ReadWriter interface {
// 	Reader
// 	Writer
// }

// type Document struct {
// 	content string
// }

// func (d Document) Read() string {
// 	return d.content
// }

// func (d *Document) Write(content string) {
// 	d.content = content
// }

// func (a author) details() {
// 	fmt.Printf("Author Name: %s", a.a_name)
// 	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
// 	fmt.Printf("\nCollege Name: %s", a.college)
// 	fmt.Printf("\nSalary: %d", a.salary)
// 	fmt.Printf("\nPublished articles: %d", a.particles)
// }

// // Implementing method
// // of the interface 2
// func (a author) articles() {
// 	pendingarticles := a.tarticles - a.particles
// 	fmt.Printf("\nPending articles: %d", pendingarticles)
// }

// // Structure
// type author struct {
// 	a_name    string
// 	branch    string
// 	college   string
// 	year      int
// 	salary    int
// 	particles int
// 	tarticles int
// }

// type AuthorDetails interface {
// 	details()
// }

// // Interface 2
// type AuthorArticles interface {
// 	articles()
// }

// type Circle struct {
// 	radius float64
// }

// type Rectangle struct {
// 	length, width float64
// }

// func (c Circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (r Rectangle) area() float64 {
// 	return r.length * r.width
// }

// func printArea(s interface{}) {
// 	value, ok := s.(Shape)
// 	if ok {
// 		fmt.Println("Area:", value.Area())
// 	} else {
// 		fmt.Println("Not a Shape interface")
// 	}
// }

// type Shape interface {
// 	Area() float64
// }
// type Square struct {
// 	side float64
// }

// func (s Square) Area() float64 {
// 	return s.side * s.side
// }

// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Circle struct {
// 	radius float64
// }

// type Rectangle struct {
// 	length, width float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }
// func (r Rectangle) Area() float64 {
// 	return r.length * r.width
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.length + r.width)
// }
