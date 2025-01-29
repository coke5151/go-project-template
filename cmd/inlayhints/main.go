package main

import "fmt"

type Gender int

// Inlay Hints: Constant Values
const (
	Male Gender = iota
	Female
	Other
)

type Person struct {
	Name   string
	Age    int
	Gender Gender
}

func Add(a, b int) int {
	return a + b
}

func main() {
	// Inlay Hints: Assign Variable Types
	first, second := 1, 2
	fmt.Println(first, second)

	// Inlay Hints: Composite Literal Fields
	data := []struct {
		id   int
		text string
	}{
		{1, "Hello"}, // Inlay Hints: Composite Literal Fields
		{2, "World"},
	}
	fmt.Println(data)
	fmt.Println("1 + 2 =", Add(1, 2)) // Inlay Hints: Parameter Names

	// Inlay Hints: Range Variable Types (slice for eg)
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers { // Inlay Hints: Range Variable Types
		fmt.Printf("%d: %d\n", i, num)
	}

	// Inlay Hints: Range Variable Types (map for eg)
	scores := map[string]int{
		"Alice":   85,
		"Bob":     92,
		"Charlie": 78,
	}
	for name, score := range scores { // Inlay Hints: Range Variable Types
		fmt.Printf("%s: %d\n", name, score)
	}
}
