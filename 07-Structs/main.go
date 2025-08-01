package main

import "fmt"

func main() {
	fmt.Println("Welcome to class 07")
	abdullah := User{"Abdullah", "abdullah@dev.com", true, 17}

	fmt.Printf("abdullah's details are: %+v\n", abdullah)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
