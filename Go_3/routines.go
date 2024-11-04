package main

import (
	"fmt"
	"time"
)

func somePrint(n string) {
	fmt.Println(n)
}

func main() {
	go somePrint("1")
	go somePrint("2")
	go somePrint("3")
	go somePrint("4")
	go somePrint("5")

	time.Sleep(100)
	fmt.Println("hello from the main function")
}

// go routines are used to call the function asynchronously
// it forked the function from the main and do it
// main function won't wait to complete all the go routines
