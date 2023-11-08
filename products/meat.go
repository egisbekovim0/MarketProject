package products

// MeatFactory is a factory for meat products
type MeatFactory struct{}

// CreateProduct creates a meat product
func (mf MeatFactory) CreateProducts() []Product {
	return []Product{
		&Meat{Name: "Beef", Price: 10.0},
		&Meat{Name: "Chicken", Price: 7.0},
		&Meat{Name: "Lamb", Price: 14.0},
		&Meat{Name: "Steak", Price: 11.0},
	}
}

// Meat represents a meat product
type Meat struct {
	Name  string
	Price float64
}

// GetName returns the name of the meat product
func (m *Meat) GetName() string {
	return m.Name
}

// GetPrice returns the price of the meat product
func (m *Meat) GetPrice() float64 {
	return m.Price
}
