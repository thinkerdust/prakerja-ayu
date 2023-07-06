package main

import "fmt"

func hitungLuasTrapesium(alas, tinggi, sisiKiri, sisiKanan float64) float64 {
	luas := (alas + sisiKiri) * tinggi / 2
	return luas
}

func main() {
	var alas, tinggi, sisiKiri, sisiKanan float64

	fmt.Print("Masukkan panjang alas: ")
	fmt.Scan(&alas)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)

	fmt.Print("Masukkan panjang sisi kiri: ")
	fmt.Scan(&sisiKiri)

	fmt.Print("Masukkan panjang sisi kanan: ")
	fmt.Scan(&sisiKanan)

	luas := hitungLuasTrapesium(alas, tinggi, sisiKiri, sisiKanan)
	fmt.Println("Luas trapesium:", luas)
}
