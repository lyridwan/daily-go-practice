package main

import "fmt"

func main() {
	// strong type
	var firstName string = "Muhammad"

	var lastName string
	lastName = "Ridwan"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo-hardType", firstName, lastName+"!")

	// soft type
	stFirstName := "Muhammad"
	stLastName := "Ridwan"
	fmt.Printf("halo-softType %s %s!\n", stFirstName, stLastName)

	// multi variable declaration
	var one, two, three string
	one, two, three = "satu", "dua", "tiga"
	fmt.Printf("multi variable declaration %s %s %s\n", one, two, three)

	var four, five, six string = "empat", "lima", "enam"
	fmt.Printf("multi variable declaration %s %s %s\n", four, five, six)

	seven, eight, nine := "tujuh", "delapan", "sembilan"
	fmt.Printf("multi variable declaration %s %s %s\n", seven, eight, nine)

	oneInt, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Printf("multi variable declaration %d %t %f %s\n", oneInt, isFriday, twoPointTwo, say)

	// pointer
	namePointer := new(string)
	fmt.Println("pointer address", namePointer)
	fmt.Println("pointer value", *namePointer)

}
