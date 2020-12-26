package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	"github.com/peacefixation/projecteuler/problem"
)

func main() {

	p := flag.Int("problem", 0, "the problem to solve")
	list := flag.Bool("list", false, "list the problems that have been solved by this program")
	flag.Parse()

	if *list {
		printProblems()
		return
	}

	if *p == 0 {
		flag.Usage()
		return
	}

	err := runProblem(*p)
	if err != nil {
		log.Fatal(err)
	}
}

func runProblem(n int) error {
	m := problem.Map()

	if p, ok := m[n]; ok {
		fmt.Println(p.Name())
		fmt.Println()
		fmt.Println(p.Description())
		fmt.Println()
		fmt.Println(p.Run())
	} else {
		return fmt.Errorf("Invalid problem: %d", n)
	}

	return nil
}

func printProblems() {
	m := problem.Map()
	keys := make([]int, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%3d  %s\n", k, m[k].Name())
	}
}
