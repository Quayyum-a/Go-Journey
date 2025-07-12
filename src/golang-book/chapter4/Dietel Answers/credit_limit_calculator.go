package main

import "fmt"

func credit_limit_calculator() {
	var accountNumber int
	var beginningBalance int
	var totalCharges int
	var totalCredits int
	var creditLimit int

	fmt.Print("Enter account number: ")
	fmt.Scan(&accountNumber)

	fmt.Print("Enter balance at the beginning of the month: ")
	fmt.Scan(&beginningBalance)

	fmt.Print("Enter total of all items charged this month: ")
	fmt.Scan(&totalCharges)

	fmt.Print("Enter total of all credits applied this month: ")
	fmt.Scan(&totalCredits)

	fmt.Print("Enter allowed credit limit: ")
	fmt.Scan(&creditLimit)

	newBalance := beginningBalance + totalCharges - totalCredits

	fmt.Printf("\nAccount: %d\n", accountNumber)
	fmt.Printf("Credit limit: %d\n", creditLimit)
	fmt.Printf("New balance: %d\n", newBalance)

	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded.")
	} else {
		fmt.Println("Credit limit not exceeded.")
	}
}

func main() {
	credit_limit_calculator()
}