package main

import "fmt"

func main() {
	//a := "Hello world"
	//
	//p := &a
	//fmt.Println(p)
	//fmt.Println(*p)
	//*p = "Oh my"
	//fmt.Println(a)

	var test interface{} = "help!"
	switch t := test.(type) {
	case string:
		fmt.Println("Variable is of type string!")
	case int:
		fmt.Println("Variable is of type int!")
	case float32:
		fmt.Println("Variable is of type float32!")
	default:
		fmt.Println("Variable type unknown")
		fmt.Println(t)
	}
	//fmt.Println(t)
}
