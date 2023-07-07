package main

import (
	"fmt"
)

func countOccurrences(slice []string, target string) int {
	count := 0
	for _, str := range slice {
		if str == target {
			count++
		}
	}
	return count
}

func main() {
	data := []string{"apple", "banana", "orange", "apple", "grape", "apple"}
	target := "apple"

	occurrences := countOccurrences(data, target)
	fmt.Printf("Jumlah kemunculan string \"%s\" dalam slice: %d\n", target, occurrences)
}
