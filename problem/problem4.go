package problem

import (
	"fmt"

	"github.com/peacefixation/projecteuler/util"
)

// Problem4 Project Euler problem #4
type Problem4 struct{}

func init() {
	addProblemToMap(4, Problem4{})
}

// Name the name of the problem
func (p Problem4) Name() string {
	return "Largest palindrome product"
}

// Description the description of the problem
func (p Problem4) Description() string {
	return `A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
	
Find the largest palindrome made from the product of two 3-digit numbers.`
}

// Run the solution to the problem and print the result
func (p Problem4) Run() string {
	largest := -1

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if util.IsPalindrome(fmt.Sprintf("%d", product)) && product > largest {
				largest = product
			}
		}
	}

	return fmt.Sprintf("%d", largest)
}
