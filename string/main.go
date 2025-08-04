package main

import (
	"fmt"
	"strings"
)

func main() {
	// ! String literal
	// str := "Chi bao"
	// str2 := `Ly Minh Thang`

	// fmt.Println(str)
	// fmt.Println(str2)

	// fmt.Printf("\nCharacter = %c", str[2])

	// ! Create a string from a slice
	// mybytes := []byte{'a', 'b', 'c'}
	// mystr := string(mybytes)
	// fmt.Println(mystr)

	// ! Length of string
	// mystr := "Welcome to GeeksforGeeks ??????"
	// fmt.Println(len(mystr))
	// fmt.Println(utf8.RuneCountInString(mystr))

	// ! Compare string
	// s1 := "Hello"
	// s2 := "Geeks"
	// s3 := "Hello"

	// // Equality and inequality
	// fmt.Println("s1 == s2:", s1 == s2) // false
	// fmt.Println("s1 == s3:", s1 == s3) // true
	// fmt.Println("s1 != s2:", s1 != s2) // true

	// // Lexicographical comparison
	// fmt.Println("s1 < s2:", s1 < s2)   // true
	// fmt.Println("s1 > s2:", s1 > s2)   // false
	// fmt.Println("s1 <= s3:", s1 <= s3) // true
	// fmt.Println("s1 >= s3:", s1 >= s3) // true

	// ! concatenate
	// s1 := "Hello, "
	// s2 := "Geeks!"

	// // Concatenating using + operator
	// result := s1 + s2
	// fmt.Println("", result)

	// s1 := "Hello, "
	// s2 := "Geeks!"
	// var b bytes.Buffer
	// b.WriteString(s1)
	// b.WriteString(s2)
	// fmt.Println("", b.String())

	// s1 := "Hello, "
	// s2 := "Geeks!"
	// result := fmt.Sprintf("%s %s", s1, s2)
	// fmt.Println("", result)

	// s1 := "Hello, "
	// s2 := "Geeks!"
	// result := strings.Join([]string{s1, s2}, "")
	// fmt.Println(result)

	// s1 := "Hello, "
	// s2 := "Geeks!"
	// var builder strings.Builder
	// builder.WriteString(s1)
	// builder.WriteString(s2)
	// fmt.Println(builder.String())

	// ! Trim string
	// s := "@@Hello,@@ Geeks!!"
	// result := strings.Trim(s, "@!")
	// fmt.Println(result)

	//  s := "   Hello, Geeks   "
	// result := strings.TrimSpace(s)
	// fmt.Println(result)

	// s := "@@Hello, Geeks!!"
	// result := strings.TrimPrefix(s, "@@H")
	// fmt.Println(result)

	// ! Split string
	// s := "Welcome,to,GeeksforGeeks"
	// fmt.Println("", s)

	// result := strings.Split(s, ",")
	// for _, str := range result {
	// 	fmt.Print(str, " ")
	// }

	// s := "Welcome,to,GeeksforGeeks"
	// fmt.Println("", s)

	// result := strings.SplitAfter(s, "")
	// // fmt.Println("Result:", result)
	// for _, str := range result {
	// 	fmt.Println(str, " ")
	// }

	// s := "Welcome,to,GeeksforGeeks"
	// fmt.Println("", s)

	// result := strings.SplitN(s, ",", -2)
	// for _, str := range result {
	// 	fmt.Println(str, " ")
	// }

	// ! Check if the given characters is present
	// str1 := "Welcome to Geeks for Geeks"
	// str2 := "Here! we learn about go strings"

	// fmt.Println("Original strings")
	// fmt.Println("String 1: ", str1)
	// fmt.Println("String 2: ", str2)

	// // Checking the string present or not
	// //  Using Contains() function
	// res1 := strings.Contains(str1, "Geeks")
	// res2 := strings.Contains(str2, "GFG")

	// // Displaying the result
	// fmt.Println("\nResult 1: ", res1)
	// fmt.Println("Result 2: ", res2)

	// Creating and initializing strings
	// str1 := "Welcome to Geeks for Geeks"
	// str2 := "Here! we learn about go strings"

	// // Checking the string present or not
	// // Using ContainsAny() function
	// res1 := strings.ContainsAny(str1, "Geeks")
	// res2 := strings.ContainsAny(str2, "GFG")
	// res3 := strings.ContainsAny("GeeksforGeeks", "G & f")
	// res4 := strings.ContainsAny("GeeksforGeeks", "u | e")
	// res5 := strings.ContainsAny(" ", " ")
	// res6 := strings.ContainsAny("GeeksforGeeks", " ")

	// // Displaying the result
	// fmt.Println("\nResult 1: ", res1)
	// fmt.Println("Result 2: ", res2)
	// fmt.Println("Result 3: ", res3)
	// fmt.Println("Result 4: ", res4)
	// fmt.Println("Result 5: ", res5)
	// fmt.Println("Result 6: ", res6)

	// ! How to find the index value of specified string
	// Creating and initializing the strings
	str1 := "Welcome to the online portal of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Finding the index value of the given strings
	// Using Index() function
	res1 := strings.Index(str1, "Geeks")
	res2 := strings.Index(str2, "do")
	res3 := strings.Index(str3, "chess")
	res4 := strings.Index("GeeksforGeeks, geeks", "ks")

	// Displaying the result
	fmt.Println("\nIndex values:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)

}
