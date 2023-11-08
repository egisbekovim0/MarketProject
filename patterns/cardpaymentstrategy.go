package patterns

// CardPaymentStrategy is a payment strategy using a card
type CardPaymentStrategy struct{}

// Pay implements payment using a card
func (cs CardPaymentStrategy) Pay(total float64, availableFunds float64) bool {
	return availableFunds >= total
}
