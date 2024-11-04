package main

import "fmt"

func main() {

	myChannel := make(chan string)
	myAnotherChannel := make(chan string)

	go func() {
		myChannel <- "hello from myChannel"
	}()

	go func() {
		myAnotherChannel <- "hello from myAnotherChannel"
	}()

	select {
	case messaeFromMyChannel := <-myChannel:
		fmt.Println(messaeFromMyChannel)
	case messageFromAnotherChannel := <-myAnotherChannel:
		fmt.Println(messageFromAnotherChannel)
	}
}
