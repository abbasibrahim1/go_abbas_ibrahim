package main

import (
	"fmt"
	
)

func main() {
	// Given values
	radius := 5.0 // Replace with your desired radius
	height := 10.0 // Replace with your desired height

	// Calculate the volume
	volume := math.Pi * math.Pow(radius, 2) * height

	fmt.Printf("Volume of the tube: %.2f\n", volume)
}
