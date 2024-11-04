package main

import "fmt"

// A function that takes an integer and returns its square
// func outerF(x int) func() int {
// 	s:=10
// 	func innerF() int{
// 		return s+x
// 	}
// 	return innerF
// }

func main() {
    f := outerF 
    result := f(5) 
    fmt.Println(result()) 
}

// A function that takes an integer and returns a function
func outerF(x int) func() int { // Changed void to func() int
    s := 10
    return func() int { // Return the inner function
        return s + x
    }
}

// func main() {
//     f := outerF(5) // Call outerF with 5
//     fmt.Println(f()) // Call the returned function and print the result
// }