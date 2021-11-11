package main

import "fmt"

/*
   win  = 3 points
   draw = 1 points
   loss = 0 points
*/

func main() {
	var input string
	fmt.Scanln(&input)

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	results = append(results, input)

	fmt.Println(scanMatchResults(results...))
}

func scanMatchResults(results ...string) int {
	total := 0
	for _, char := range results {
		total += determinePoints(char)
	}
	return total
}

func determinePoints(char string) int {
	if char == "w" {
		return 3
	} else if char == "d" {
		return 1
	}
	return 0
}
