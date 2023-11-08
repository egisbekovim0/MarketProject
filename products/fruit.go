package products

// FruitFactory is a factory for fruit products
type FruitFactory struct{}

// CreateProduct creates a fruit product
func (ff FruitFactory) CreateProducts() []Product {
	return []Product{
		&Fruit{Name: "Apple", Price: 1.0},
		&Fruit{Name: "Banana", Price: 3.0},
		&Fruit{Name: "Melon", Price: 7.0},
		&Fruit{Name: "Grapes", Price: 5.0},
	}
}

// Fruit represents a fruit product
type Fruit struct {
	Name  string
	Price float64
}

// GetName returns the name of the fruit product
func (f *Fruit) GetName() string {
	return f.Name
}

// GetPrice returns the price of the fruit product
func (f *Fruit) GetPrice() float64 {
	return f.Price
}
