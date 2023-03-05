package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) EstimateDistance() float64 {
	fuelEfficiency := 1.5 // perhitungan konstanta bahan bakar per mil
	distance := c.FuelIn * fuelEfficiency
	return distance
}

func main() {
	myCar := Car{"sedan", 13.0}
	distance := myCar.EstimateDistance()
	fmt.Printf("Dengan bahan bakar sebesar %.2f liter, mobil %s dapat menempuh jarak sekitar %.2f mil.\n", myCar.FuelIn, myCar.Type, distance)
}
