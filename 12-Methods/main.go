package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")
	// no inheritance in golang; No super or parent

	nirob := User{"Nirob", "nirob@go.dev", true, 16}
	fmt.Println(nirob)
	fmt.Printf("nirob details are: %+v\n", nirob)
	fmt.Printf("Name is %v and email is %v.\n", nirob.Name, nirob.Email)
	nirob.GetStatus()
	nirob.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", nirob.Name, nirob.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}