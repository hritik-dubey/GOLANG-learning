package main

import "fmt"

// Package-scoped variable
var globalVar int = 100

func main() {
	// Function-scoped variable
	var localVar int = 10
	fmt.Println("Global Variable:", globalVar)
	fmt.Println("Local Variable:", localVar)

	// Short variable declaration
	x := 5
	fmt.Println("Short Declaration:", x)

	// Block scope
	if true {
		var blockVar int = 20
		fmt.Println("Block Variable:", blockVar)
	}
	fmt.Println("Value of Pi:", Pi)
	fmt.Println("Value of Gravity:", Gravity)
	fmt.Println("Programming Language:", Language)
	fmt.Println("Days of the Week:")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	// fmt.Println(blockVar) // Uncommenting this line will cause an error
}

const (
	Pi       = 3.14 // Constant without explicit type
	Gravity  = 9.81 // Constant without explicit type
	Language = "Go" // Constant without explicit type
)

const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)
