package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	// Constants
	const pi = 3.14159
	const projectName = "My First Go project!"

	// Variables
	var a int = 10
	var b float64 = 3.14
	var name string = "nolan"
	var isInitialised bool = true

	// Short variable declarations
	i := 100
	j := 2.8234

	// Struct initialization
	var nolan Person
	nolan.name = "nolan"

	// Slice initialization
	var m []int = []int{1, 2, 3}
	n := []string{"Go", "Python", "Java"}

	// Print variables and their types
	fmt.Printf("a: %d, type: %T \n", a, a)
	fmt.Printf("b: %f, type: %T \n", b, b)
	fmt.Printf("name: %s, type: %T \n", name, name)
	fmt.Printf("isInitialised: %t, type: %T \n", isInitialised, isInitialised)
	fmt.Printf("i: %d, type: %T \n", i, i)
	fmt.Printf("j: %f, type: %T \n", j, j)
	fmt.Printf("nolan: %+v, type: %T \n", nolan)
	fmt.Printf("m: %v, type: %T \n", m, m)
	fmt.Printf("n: %v, type: %T \n", n, n)
}
