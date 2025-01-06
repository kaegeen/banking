package main

import (
	"fmt"
)

func main() {
	balance := 0.0

	fmt.Println("Simple Banking System")
	fmt.Println("=====================")
	fmt.Println("Commands:")
	fmt.Println("1. deposit <amount> - Add money to your account.")
	fmt.Println("2. withdraw <amount> - Remove money from your account.")
	fmt.Println("3. check - Check your current balance.")
	fmt.Println("4. exit - Exit the program.\n")

	for {
		fmt.Print("Enter command: ")
		var command string
		var amount float64
		_, err := fmt.Scan(&command, &amount)

		if command == "exit" {
			fmt.Println("Thank you for using the banking system. Goodbye!")
			break
		}

		if err != nil && command != "check" {
			fmt.Println("Invalid command. Please try again.\n")
			continue
		}

		switch command {
		case "deposit":
			if amount <= 0 {
				fmt.Println("Deposit amount must be positive.\n")
			} else {
				balance += amount
				fmt.Printf("Deposited %.2f. New balance: %.2f\n\n", amount, balance)
			}
		case "withdraw":
			if amount <= 0 {
				fmt.Println("Withdrawal amount must be positive.\n")
			} else if amount > balance {
				fmt.Printf("Insufficient funds! Your current balance is %.2f\n\n", balance)
			} else {
				balance -= amount
				fmt.Printf("Withdrew %.2f. New balance: %.2f\n\n", amount, balance)
			}
		case "check":
			fmt.Printf("Your current balance is: %.2f\n\n", balance)
		default:
			fmt.Println("Unknown command. Please use deposit, withdraw, check, or exit.\n")
		}
	}
}
