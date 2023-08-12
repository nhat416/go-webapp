package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	/* maps
	myMap := make(map[string]User)

	me := User{
		FirstName: "Go",
		LastName:  "Leafs",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	//var myNewVar float32
	//myNewVar = 11.1
	*/

	var myString string
	myString = "Fish"

	log.Println(myString)

	var mySlice []string
	var mySliceTwo []int

	mySlice = append(mySlice, "Toronto")
	mySlice = append(mySlice, "Maple")
	mySlice = append(mySlice, "Leafs")

	mySliceTwo = append(mySliceTwo, 2)
	mySliceTwo = append(mySliceTwo, 1)
	mySliceTwo = append(mySliceTwo, 3)

	log.Println(mySlice)
	log.Println(mySliceTwo)

	sort.Ints(mySliceTwo)
	log.Println(mySliceTwo)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[0:2])

	names := []string{"Tavares", "Matthews", "Marner", "Nylander", "Rielly"}
	log.Println(names)
}
