package main

import "fmt"

func main() {
	fmt.Println("pointers")

	var ptr *int

	fmt.Println(ptr)

	myNumber := 23

	var myNumberPointer *int

	myNumberPointer = &myNumber

	*myNumberPointer += 20

	fmt.Println(myNumberPointer, myNumber)

}
