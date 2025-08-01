package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Present time is")
	presnetTime := time.Now()
	fmt.Println(presnetTime.Format("02-01-2006 15:03:04 Monday"))
}