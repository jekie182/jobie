package main

import (
	"fmt"


	"https://github.com/jekie182/jobie/tree/master/api/controllers"
)

func main() {
	fmt.Println("app works")
	defer controllers.ControllerAccesTetst()

	first()
	//panic(1)
	//defer second()

	third()

	//http.Handle("/", Test)
}

func first() {
	fmt.Println("first function has been done")
}

func third() {
	fmt.Println("third function has been done")
}
