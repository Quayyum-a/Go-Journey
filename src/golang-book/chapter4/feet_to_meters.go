package main

import (
	"fmt"
)

func main() {
	var feet float64
	fmt.Print("Enter length in feet: ")
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	fmt.Printf("%.2f feet is %.4f meters\n", feet, meters)
} 