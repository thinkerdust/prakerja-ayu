package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	// Memeriksa divisibilitas hingga akar kuadrat dari bilangan
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var num int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Println("Bilangan prima")
	} else {
		fmt.Println("Bukan bilangan prima")
	}
}
