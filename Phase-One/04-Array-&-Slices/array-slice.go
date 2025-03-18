package main

import "fmt"

func main() {
	// Array - fixed size
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers)

	// Accessing elements
	fmt.Printf("Accessing element One: %d\n", numbers[0])

	// slice - dynamic size
	var cities = []string{"New York", "Paris", "London"}
	fmt.Println(cities)

	// adding elements
	cities = append(cities, "Tokyo")
	fmt.Println(cities)

	// removing elements
	cities = append(cities[:2], cities[3:]...)
	fmt.Println(cities)

	// copying slices
	var newCities = make([]string, len(cities))
	copy(newCities, cities)
	fmt.Println(newCities)

	// iterating over slices
	for i, city := range cities {
		fmt.Printf("City %d: %s\n", i, city)
	}

	// multidimensional slices
	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	
	
	
}	