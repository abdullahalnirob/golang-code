package main
import "fmt"

func main()  {
	fmt.Println("maps in golang")

	languages :=make(map[string]string)
	languages["JS"]="JavaScript"
	languages["GO"]="Golang"
	languages["PY"]="Python"

	fmt.Println(languages["GO"]) //Golang
	fmt.Println(languages["PY"]) //Python
}