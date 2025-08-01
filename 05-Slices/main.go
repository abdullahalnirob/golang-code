package main
import "fmt"

func main()  {
	fmt.Println("Welcome to slices")
	var vegList = []string{"Tomato", "Potatu","Letus","Cabbege","Bean"}
	
	fmt.Println(vegList[1:3])
	fmt.Println(vegList[0:2])
}