package main

import (
	"MarketProject/patterns"
	"MarketProject/products"
	"fmt"
)

func main() {
	fmt.Printf("Welcome to the Marketplace!\n")
	fmt.Println("Choose a category:")
	fmt.Println("1. Meat")
	fmt.Println("2. Fruit")

	var categoryChoice int
	fmt.Print("Enter your choice (1 for Meat, 2 for Fruit): ")
	fmt.Scanln(&categoryChoice)

	var productFactory products.Factory
	var categoryName string

	switch categoryChoice {
	case 1:
		productFactory = &products.MeatFactory{}
		categoryName = "Meat"
	case 2:
		productFactory = &products.FruitFactory{}
		categoryName = "Fruit"
	default:
		fmt.Println("Invalid choice")
		return
	}

	productsList := productFactory.CreateProducts()

	//fmt.Printf("Choose a %s product:\n", categoryName)
	for i, product := range productsList {
		fmt.Printf("%d. %s - $%.2f\n", i+1, product.GetName(), product.GetPrice())
	}

	var productChoice int
	fmt.Printf("Enter your choice from %s \n", categoryName)
	fmt.Scanln(&productChoice)

	if productChoice < 1 || productChoice > len(productsList) {
		fmt.Println("Invalid product choice")
		return
	}

	cart := []patterns.CartItem{}
	cashPaymentStrategy := &patterns.CashPaymentStrategy{}
	cardPaymentStrategy := &patterns.CardPaymentStrategy{}

	funds := patterns.GetFundsInstance()
	availableFunds := funds.Amount

	observers := []patterns.Observer{&CartObserver{}}

	fmt.Printf("You have $%.2f in available funds.\n", availableFunds)

	selectedProduct := productsList[productChoice-1]

	cartItem := patterns.CartItem{
		Name:  selectedProduct.GetName(),
		Price: selectedProduct.GetPrice(),
	}
	cart = append(cart, cartItem)
	observers[0].Update(selectedProduct.GetName(), selectedProduct.GetPrice())

	for {
		var buyChoice int
		fmt.Println("Do you want to buy another one? 1 = yes, 2 = no")
		fmt.Scanln(&buyChoice)

		if buyChoice == 1 {
			// User wants to buy another product
			fmt.Println("Choose another product:")
			// Repeat the product selection logic and add the selected product to the cart
			// Use the same logic you've already written for product selection and adding to the cart
			productsList := productFactory.CreateProducts()
			for i, product := range productsList {
				fmt.Printf("%d. %s - $%.2f\n", i+1, product.GetName(), product.GetPrice())
			}

			var productChoice int
			fmt.Printf("Enter your choice: ")
			fmt.Scanln(&productChoice)

			if productChoice < 1 || productChoice > len(productsList) {
				fmt.Println("Invalid product choice")
				return
			}

			selectedProduct := productsList[productChoice-1]
			cart = append(cart, patterns.CartItem{Name: selectedProduct.GetName(), Price: selectedProduct.GetPrice()})

		} else if buyChoice == 2 {
			// User wants to proceed to payment

			// Proceed to payment logic (use the existing payment logic)
			total := 0.0
			for _, item := range cart {
				total += item.Price
			}

			paymentChoice := 0
			fmt.Printf("Select a payment method:\n1. Cash\n2. Card\nEnter your choice: ")
			fmt.Scanln(&paymentChoice)

			var paymentStrategy patterns.PaymentStrategy
			switch paymentChoice {
			case 1:
				paymentStrategy = cashPaymentStrategy
			case 2:
				paymentStrategy = cardPaymentStrategy
			default:
				fmt.Println("Invalid payment choice")
				return
			}

			if paymentStrategy.Pay(total, availableFunds) {
				fmt.Println("Payment successful. Enjoy your purchase!")
				funds.Amount -= total
				fmt.Printf("Remaining funds: $%.2f\n", funds.Amount)
			} else {
				fmt.Println("Insufficient funds. Payment failed.")
			}
			// After payment, you can exit the loop
			break

		} else {
			fmt.Println("Invalid choice")
		}
	}

}

// CartObserver is an observer for the cart
type CartObserver struct{}

// Update notifies the observer about the selected product
func (co *CartObserver) Update(productName string, productPrice float64) {
	fmt.Printf("Added %s - $%.2f to the cart.\n", productName, productPrice)
}
