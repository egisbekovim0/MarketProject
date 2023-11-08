package patterns

type Observer interface {
	Update(productName string, productPrice float64)
}

// Cart is the subject being observed
type Cart struct {
	Observers []Observer
	Products  []CartItem // Add a slice to store the products in the cart
}

// CartItem represents an item in the cart
type CartItem struct {
	Name  string
	Price float64
}

// Attach adds an observer to the cart
func (c *Cart) Attach(observer Observer) {
	c.Observers = append(c.Observers, observer)
}

// AddProduct adds a product to the cart and notifies observers
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
