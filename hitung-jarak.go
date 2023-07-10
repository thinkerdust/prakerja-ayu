package main

import (
	"fmt"
)

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) CalculateDistance() float64 {
	fuelConsumptionRate := 1.5 // konsumsi bahan bakar per mil
	distance := c.FuelIn / fuelConsumptionRate
	return distance
}

func main() {
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 10.0, // jumlah bahan bakar dalam liter
	}

	distance := myCar.CalculateDistance()
	fmt.Printf("Perkiraan jarak yang bisa ditempuh: %.2f mil\n", distance)
}
