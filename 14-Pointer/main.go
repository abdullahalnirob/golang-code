package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointer")

	MyNumber := 23 //Assign a number

	ptr := &MyNumber //Save memory location into ptr
	fmt.Println("Value of actual memory location", ptr)  //0xc00000a0c8
	fmt.Println("Value of actual pointers", *ptr)  //23

}
