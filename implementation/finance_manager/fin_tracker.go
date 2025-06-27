/*
Chapter 1: Foundation Setup

	Display a welcome message and menu options
	Accept user input for different operations (add income, add expense, view summary, quit)
	Use functions to handle each menu option
	Implement input validation using control flow

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("*** Welcome to FinTrack ***")

	userProfile := map[string]map[string]float64{
		"rohan": map[string]float64{
			"income":  0,
			"expense": 0,
		},
	}

	transactions := []map[string]string{}

	menu_driven_app(&userProfile, &transactions)
	// income, expense := calcTransaction(&transactions)

}

func menu_driven_app(profile *map[string]map[string]float64, transaction *[]map[string]string) {

	for i := 1; i > 0; i++ {
		choice := 0
		username := "rohan"

		fmt.Println(strings.Repeat("_", 80))
		fmt.Println("App Used: ", i)
		fmt.Println(strings.Repeat("-", 80))

		fmt.Println(strings.Repeat("-", 80))
		fmt.Println("Total Income: ", (*profile)[username]["income"])
		fmt.Println("Total Expenses: ", (*profile)[username]["expense"])
		fmt.Println(strings.Repeat("-", 80))

		fmt.Println("")
		fmt.Println("1) Add Income")
		fmt.Println("2) Add Expense")
		fmt.Println("3) View Summary")
		fmt.Println("4) Analysis and Reporting")
		fmt.Println("5) Quit")
		fmt.Println("")

		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			income := addIncome()
			*transaction = append(*transaction, map[string]string{
				"type":   "income",
				"amount": income,
			})

			_, amount := calcTransaction(&transaction)

			for uname, value := range *profile {
				if uname == "rohan" {
					value["income"] = amount
				}
			}

		case 2:
			amount, category := addExpense()
			*transaction = append(*transaction, map[string]string{
				"type":     "expense",
				"category": category,
				"amount":   amount,
			})

			expense, _ := calcTransaction(&transaction)

			// (*profile)[username]["expense"] += expense

			for uname, value := range *profile {
				if uname == "rohan" {
					value["expense"] = expense
				}
			}

		case 3:
			for uname, v := range *profile {
				if uname == username {
					income := v["income"]
					expense := v["expense"]

					diff := income - expense
					fmt.Printf("Username: %s\nIncome: %.2f\nExpense: %.2f\nSavings: %.2f\n", username, income, expense, diff)
					if diff < 0 {
						fmt.Println("You're spending too much. Try to save some money")
					}
				}
			}

		case 4:
			totalExpense, totalIncome := calcTransaction(&transaction)
			displayTransactionSummary(&transaction, totalExpense, totalIncome)

		case 5:
			fmt.Println("Quitting Tracker !!")
			return

		default:
			fmt.Println("Invalid Choice")
			return

		}
	}
}

// helper funcs
func addIncome() string {
	income := ""
	fmt.Println("Enter your income: ")
	fmt.Scanln(&income)
	return income
}

func addExpense() (string, string) {
	amount := ""
	var category string
	fmt.Printf("Enter your expense and category (e.g., 50 food): ")
	fmt.Scanln(&amount, &category)
	return amount, category
}

func calcTransaction(transactions **[]map[string]string) (float64, float64) {
	totalIncome, totalExpense := 0.0, 0.0
	for _, value := range **transactions {
		if value["type"] == "income" {
			amount, _ := strconv.ParseFloat(value["amount"], 64)
			totalIncome += amount
		}

		if value["type"] == "expense" {
			amount, _ := strconv.ParseFloat(value["amount"], 64)
			totalExpense += amount
		}
	}
	return totalExpense, totalIncome
}

func displayTransactionSummary(transaction **[]map[string]string, totalExpense, totalIncome float64) {

	if len(**transaction) == 0 {
		fmt.Println("No transaction history found")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("                          TRANSACTION HISTORY")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Printf("%-5s %-12s %-15s %-12s\n", "ID", "TYPE", "CATEGORY", "AMOUNT")
	fmt.Println(strings.Repeat("-", 80))

	for i, transaction := range **transaction {
		id := fmt.Sprintf("#%d", i+1)
		transType := strings.ToUpper(transaction["type"])
		category := strings.ToUpper(transaction["category"])
		if category == "" {
			category = "N/A"
		}
		amount := transaction["amount"]

		if amount != "" {
			if amountFloat, err := strconv.ParseFloat(amount, 64); err == nil {
				amount = fmt.Sprintf("$%.2f", amountFloat)
			}
		}

		fmt.Printf("%-5s %-12s %-15s %-12s\n", id, transType, category, amount)

	}
	fmt.Println(strings.Repeat("-", 80))
	diff := totalIncome - totalExpense

	fmt.Printf("INCOME: %.2f\n", totalIncome)
	fmt.Printf("EXPENSE: %.2f\n", totalExpense)
	fmt.Printf("SAVINGS: %.2f\n", diff)

	if diff < 0 {
		fmt.Println("Too Much Expense: Make a habit of savings !!")
	}

	fmt.Println(strings.Repeat("-", 80))

}
