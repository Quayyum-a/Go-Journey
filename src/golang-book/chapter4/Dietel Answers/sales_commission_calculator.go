package main

import "fmt"

func sales_commission_calculator(){
	
	fmt.Println(`
Item  Value
1:    239.99,
2:    129.75,
3:    99.95,
4:    350.89,`,"\n")
	itemValues := map[int]float64{
		1: 239.99,
		2: 129.75,
		3: 99.95,
		4: 350.89,
	}
	grossSales := 0.0

	fmt.Println("Enter the item numbers sold by the salesperson (1-4). Enter 0 when done:")

	for {
		var itemNumber int
		fmt.Print("Item number: ")
		fmt.Scan(&itemNumber)

		if itemNumber == 0 {
			break
		}

		itemValue, exists := itemValues[itemNumber]
		if !exists {
			fmt.Println("Invalid item number. Please enter a number between 1 and 4.")
			continue
		}

		grossSales += itemValue
	}

	basePay := 200.0
	commission := 0.09 * grossSales
	earnings := basePay + commission

	fmt.Printf("\nGross sales: $%.2f\n", grossSales)
	fmt.Printf("Earnings for the week: $%.2f\n", earnings)
}
func main(){
	sales_commission_calculator()
}