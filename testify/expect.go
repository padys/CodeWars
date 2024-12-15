package testify

import "fmt"

type TestifyExpect struct {
	solutionValue Any
}

func (r *TestifyExpect) To(expectedValue string) {
	solutionValue := Equal(r.solutionValue)

	fmt.Printf(">>> Solution: %s\n", solutionValue)
	fmt.Printf(">>> Expected: %s\n", expectedValue)
	fmt.Printf(">>> AreEqual: %t\n", solutionValue == expectedValue)

	printLine()
}

func Expect(solutionValue Any) *TestifyExpect {
	return &TestifyExpect{
		solutionValue: solutionValue,
	}
}
