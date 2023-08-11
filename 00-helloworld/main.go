package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World!")
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)

	changeUsingPointer(&myString)
	log.Println("After using changeUsingPointer, myString is set to", myString)

}

func changeUsingPointer(s *string) {
	log.Println("s is", s)
	newValue := "Red"
	*s = newValue
}
