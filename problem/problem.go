package problem

// Problem details of a Project Euler problem including the function that solves it
type Problem interface {
	Name() string
	Description() string
	Run() string
}

var m map[int]Problem

func init() {
	m = make(map[int]Problem, 0)
}

func addProblemToMap(n int, p Problem) {
	m[n] = p
}

// Map return the map of problems
func Map() map[int]Problem {
	return m
}
