package patterns

type DiscountStrategy interface {
	ApplyDiscount(total float64) float64
}
