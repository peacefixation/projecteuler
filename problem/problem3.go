package problem

import "fmt"

// Problem3 Project Euler problem #3
type Problem3 struct{}

func init() {
	addProblemToMap(3, Problem3{})
}

// Name the name of the problem
func (p Problem3) Name() string {
	return "Largest prime factor"
}

// Description the description of the problem
func (p Problem3) Description() string {
	return `The prime factors of 13195 are 5, 7, 13 and 29.
	
What is the largest prime factor of the number 600851475143 ?`
}

// Run the solution to the problem and print the result
func (p Problem3) Run() string {
	n := 600851475143
	i := 2
	for n > 1 {
		if n%i == 0 {
			n /= i
		} else {
			i++
		}
	}

	return fmt.Sprintf("%d", i)
}
