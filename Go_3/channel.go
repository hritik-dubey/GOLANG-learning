package main

import "fmt"

func main() {
	myChannel := make(chan string)
	go func() { myChannel <- "hello" }()
	msg := <-myChannel
	fmt.Println(msg)
}
