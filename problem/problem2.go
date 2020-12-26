package problem

import "fmt"

// Problem2 Project Euler problem #2
type Problem2 struct{}

func init() {
	addProblemToMap(2, Problem2{})
}

// Name the name of the problem
func (p Problem2) Name() string {
	return "Even Fibonacci numbers"
}

// Description the description of the problem
func (p Problem2) Description() string {
	return `Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.`
}

// Run the solution to the problem and print the result
func (p Problem2) Run() string {
	const max = uint64(4000000)
	sum := uint64(0)

	f1, f2 := uint64(1), uint64(1)
	for f1 <= max && f2 <= max {
		f3 := f1 + f2
		f1 = f2
		f2 = f3

		if f3%2 == 0 {
			sum += f3
		}
	}

	return fmt.Sprintf("%d", sum)
}