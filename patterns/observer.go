package patterns

import "fmt"

type Observer interface {
	Update(productName string, productPrice float64)
}
type CartObserver struct{}

type Cart struct {
	Observers []Observer
	Products  []CartItem
}

type CartItem struct {
	Name  string
	Price float64
}

func (c *Cart) Attach(observer Observer) {
	c.Observers = append(c.Observers, observer)
}

func (c *Cart) AddProduct(productName string, productPrice float64) {
	// Create a new cart item
	item := CartItem{Name: productName, Price: productPrice}

	// Add the product item to the cart
	c.Products = append(c.Products, item)

	// Notify observers
	for _, observer := range c.Observers {
		observer.Update(productName, productPrice)
	}
}

// Update notifies the observer about the selected product
func (co *CartObserver) Update(productName string, productPrice float64) {
	fmt.Printf("Added %s - $%.2f to the cart.\n", productName, productPrice)
}
