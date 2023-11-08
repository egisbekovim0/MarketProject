package patterns

type NoDiscountStrategy struct{}

func (nds *NoDiscountStrategy) ApplyDiscount(total float64) float64 {
	return total // No discount applied
}
