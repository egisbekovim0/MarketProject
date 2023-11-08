package products

type Product interface {
	GetName() string
	GetPrice() float64
}

type Factory interface {
	CreateProducts() []Product
}
