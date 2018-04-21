package main

import (
	"fmt"
)

func main() {
	fmt.Println("app works")
	//first()
	//panic(1)
	//defer second()

	//third()
	//http.Handle("/", Test)
}

func first() {
	fmt.Println("first function has been done")
}

func second() {
	/*for i := 0; i < 166; i++ {
		fmt.Print(".")
	}*/
	recover()
	fmt.Println("second function has been done")
}

func third() {
	fmt.Println("third function has been done")
}
