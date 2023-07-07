package main

import (
	"fmt"
	"strings"
)

func findUniqueNumbers(input string) []int {
	numbers := strings.Fields(input) // Memisahkan angka-angka berdasarkan spasi
	counts := make(map[string]int)

	// Menghitung kemunculan setiap angka dalam input
	for _, num := range numbers {
		counts[num]++
	}

	var uniqueNums []int

	// Menyimpan angka-angka yang hanya muncul satu kali dalam uniqueNums
	for _, num := range numbers {
		if counts[num] == 1 {
			var convertedNum int
			_, err := fmt.Sscanf(num, "%d", &convertedNum)
			if err == nil {
				uniqueNums = append(uniqueNums, convertedNum)
			}
		}
	}

	return uniqueNums
}

func main() {
	input := "1 2 3 4 5 4 3 2 6 7 8 9 9 8"
	uniqueNumbers := findUniqueNumbers(input)
	fmt.Println(uniqueNumbers)
}
