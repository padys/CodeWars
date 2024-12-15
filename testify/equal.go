package testify

import "fmt"

func Equal(expectedValue Any) string {
	return fmt.Sprint(expectedValue)
}
