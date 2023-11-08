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

type Fruit struct {
	Name  string
	Price float64
}

func (f *Fruit) GetName() string {
	return f.Name
}

func (f *Fruit) GetPrice() float64 {
	return f.Price
}
