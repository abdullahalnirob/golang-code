package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")
	// SAMPLE
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	day := []string{"Saturday", "Sunday", "Monday"}
	// METHOD 1 TO PRINT ARRAY
	for j := 0; j < len(day); j++ {
		fmt.Println(day[j])
	}
	// METHOD 2 TO PRINT ARRAY
	for p := range day {
		fmt.Println(day[p])
	}
}
