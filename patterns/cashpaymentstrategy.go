package patterns

// CashPaymentStrategy is a payment strategy using cash
type CashPaymentStrategy struct{}

// Pay implements payment using cash
func (cs CashPaymentStrategy) Pay(total float64, availableFunds float64) bool {
	return availableFunds >= total
}
