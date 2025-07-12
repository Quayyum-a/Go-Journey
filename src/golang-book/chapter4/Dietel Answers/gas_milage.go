package main

import "fmt"

func main() {
	var totalMiles, totalGallons uint64

	for {
		var miles, gallons uint64

		fmt.Print("Enter miles driven (0 to quit): ")
		fmt.Scanf("%d", &miles)
		if miles == 0 {
			break
		}

		fmt.Print("Enter gallons used (0 to quit): ")
		fmt.Scanf("%d", &gallons)
		if gallons == 0 {
			break
		}

		tripMPG := float64(miles) / float64(gallons)
		fmt.Printf("Miles per gallon for this trip: %.2f\n", tripMPG)

		totalMiles += miles
		totalGallons += gallons

		combinedMPG := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined miles per gallon: %.2f\n", combinedMPG)
	}
}