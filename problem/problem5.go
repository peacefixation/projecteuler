package problem

import "fmt"

// Problem5 Project Euler problem #5
type Problem5 struct{}

func init() {
	addProblemToMap(5, Problem5{})
}

// Name the name of the problem
func (p Problem5) Name() string {
	return "Smallest multiple"
}

// Description the description of the problem
func (p Problem5) Description() string {
	return `2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?`
}

// Run the solution to the problem and print the result
func (p Problem5) Run() string {
	for i := 2520; i < 2147483647; i++ {
		if p.divisible1to20(i) {
			return fmt.Sprintf("%d", i)
		}
	}
	return "not found"
}

func (p Problem5) divisible1to20(n int) bool {
	for j := 2; j <= 20; j++ {
		if n%j != 0 {
			return false
		}
	}
	return true
}
