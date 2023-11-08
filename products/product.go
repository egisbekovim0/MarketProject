package products

// Product represents a product
type Product interface {
	GetName() string
	GetPrice() float64
}

// Factory defines the factory to create products
type Factory interface {
	CreateProducts() []Product // New method to create multiple products
}
