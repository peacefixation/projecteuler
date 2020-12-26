package problem

import "fmt"

// Problem1 Project Euler problem #1
type Problem1 struct{}

func init() {
	addProblemToMap(1, Problem1{})
}

// Name the name of the problem
func (p Problem1) Name() string {
	return "Multiples of 3 and 5"
}

// Description the description of the problem
func (p Problem1) Description() string {
	return `If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.`
}

// Run the solution to the problem and print the result
func (p Problem1) Run() string {
	sum := 0

	for n := 0; n <= 999; n++ {
		if n%3 == 0 || n%5 == 0 {
			sum += n
		}
	}

	return fmt.Sprintf("%d", sum)
}
