package main

import "fmt"

func main() {
	var num int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&num)

	if num%7 == 0 {
		fmt.Println("Bilangan adalah kelipatan 7")
	} else {
		fmt.Println("Bilangan bukan kelipatan 7")
	}
}
