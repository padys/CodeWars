package testify

import "fmt"

func Run() {
	printLine()

	for name, fnc := range describe {
		fmt.Println("> Testify Describe:", name)
		printLine()

		fnc()
	}
}
