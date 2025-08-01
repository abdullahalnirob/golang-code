package main

import "fmt"

func main() {
	fmt.Println("Wecole to array")
	var fruiteList [3]string
	fruiteList[0] = "Apple"
	fruiteList[1] = "Banana"
	fruiteList[2] = "Mango"
	fmt.Println("fruit lists are:", fruiteList)

	// Second Method
	var vegList = [2]string{"Tomato", "Potatu"}
	fmt.Println("Vegetable lists are:", vegList)
}
