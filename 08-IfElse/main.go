package main

import "fmt"

func main() {
	fmt.Println("Hello If else coder")
	yourAge := -18
	if yourAge <= 18 {
		fmt.Println("You can't drive")
	} else if yourAge == 0 {
		fmt.Println("user not born")
	} else {
		fmt.Println("You can drive")
	}
}
