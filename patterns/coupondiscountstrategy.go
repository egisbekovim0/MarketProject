package patterns

type CouponDiscountStrategy struct {
	CouponCode string
}

func NewCouponDiscountStrategy(couponCode string) *CouponDiscountStrategy {
	return &CouponDiscountStrategy{CouponCode: couponCode}
}

func (cds *CouponDiscountStrategy) ApplyDiscount(total float64) float64 {
	if cds.CouponCode == "SAVE10" {
		// Apply a 10% discount for the "SAVE10" coupon code
		discount := 0.10 * total
		return total - discount
	}
	// No discount applied for other coupon codes
	return total
}
