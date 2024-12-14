package testify

import "fmt"

type Any = interface{}

type TestifyExpect struct {
	solutionValue Any
}

func (r *TestifyExpect) To(expectedValue string) {
	solutionValue := Equal(r.solutionValue)

	fmt.Println("- - - - - - - - - - - - - -")
	fmt.Printf("Solution: %s\n", solutionValue)
	fmt.Printf("Expected: %s\n", expectedValue)
	fmt.Printf("AreEqual: %t\n", solutionValue == expectedValue)
}

func Expect(solutionValue Any) *TestifyExpect {
	return &TestifyExpect{
		solutionValue: solutionValue,
	}
}

func Equal(expectedValue Any) string {
	return fmt.Sprint(expectedValue)
}
