package main

import "fmt"

func findMaxMin(numbers ...int) (int, int) {
	max := numbers[0]
	min := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	var inputNumbers [6]int

	fmt.Println("Masukkan 6 angka:")
	for i := 0; i < 6; i++ {
		fmt.Scan(&inputNumbers[i])
	}

	max, min := findMaxMin(inputNumbers[:]...)

	fmt.Println("Nilai maksimum:", max)
	fmt.Println("Nilai minimum:", min)
}
