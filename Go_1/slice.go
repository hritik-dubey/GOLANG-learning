
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3} // A slice of integers
    slice = append(slice, 4) // Adding an element
    fmt.Println(slice) // Output: [1 2 3 4]
}