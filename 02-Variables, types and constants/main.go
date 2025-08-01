package main

import "fmt"

func main() {
	var name string = "abdullah al nirob" //string
	fmt.Println(name)
	fmt.Printf("The name type is: %T", name) //check type

	var isLogin bool = true //boolean
	fmt.Println(isLogin)
	fmt.Printf("The isLogin type is: %T", isLogin) //check type

	var age int = 18 // integer
	fmt.Println(isLogin)
	fmt.Printf("The isLogin type is: %T \n", age) //check type

	// some bits intigers
	// 	uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	var smallNum float32 = 356.4443
	var bigNum float64 = 3434985.78584
	fmt.Println("\n", smallNum, bigNum)

	// ########################################### //
	// # short and second way to define variable # //
	// ########################################### //

	username := "dev.abdullahalnirob"
	fmt.Println(username)

	userAge := 17
	fmt.Println(userAge)

	userIsLogin := false
	fmt.Println(userIsLogin)


	// !CONSTANT VARIABLE //
	const myWebsite string = "https://abdullah-io.vercel.com"
	fmt.Println(myWebsite)
	
}
