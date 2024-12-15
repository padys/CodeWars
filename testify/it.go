package testify

import "fmt"

func It(name string, fnc Callback) {
	fmt.Println(">> Testify It:", name)
	printLine()

	fnc()
}
