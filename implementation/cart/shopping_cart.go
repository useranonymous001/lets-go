/*
	// array [fixed product lists]
	// slice [shopping cart]
	// remove/truncate slice after from last
	// copy slice [cheout receipt]
	// exit
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// created an product list
	productList := [5]string{"Acer Nitro v15", "Redmi Note 10s", "Notebook", "Pork", "Chicken"}

	// create a shopping cart for the user
	shoppingCart := []string{}
	for {

		option := 1

		fmt.Println("")
		fmt.Println(
			"1. View available products",
			"2. Add item to cart",
			"3. View cart",
			"4. Remove last item",
			"5. Checkout",
			"6. Exit",
		)

		fmt.Print("\nChoose the operation: ")
		fmt.Scanf("%v", &option)

		switch option {
		case 1:
			fmt.Println("Viewing available products")
			for i, v := range productList {
				fmt.Printf("%d -> %s\n", i+1, v)
			}

		case 2:
			input := 1
			fmt.Print("Enter item number to add in the shopping list: ")
			fmt.Scanf("%v", &input)

			if input > len(productList) || input < 0 {
				fmt.Println("Invalid product Item choosen")
				fmt.Println("Exiting Program..")
				os.Exit(1)
			}
			shoppingCart = append(shoppingCart, productList[input-1])

		case 3:
			fmt.Println("Viewing cart..")
			fmt.Println("Your cart: ", shoppingCart)
			fmt.Println(len(shoppingCart), cap(shoppingCart))

		case 4:
			fmt.Println("removing last item from cart..")
			if len(shoppingCart) > 0 {
				shoppingCart = shoppingCart[:len(shoppingCart)-1] // creating a new slice without only the last element
			} else {
				fmt.Println("Cart already empty")
			}
		case 5:
			fmt.Println("Checking out your cart....")
			checkout := make([]string, len(shoppingCart))
			copy(checkout, shoppingCart)
			fmt.Println(checkout)

		default:
			fmt.Println("Invalid Option, exiting the program")
			os.Exit(1)
		}
	}
}
