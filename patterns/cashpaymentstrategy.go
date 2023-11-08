package patterns

type CashPaymentStrategy struct{}

func (cs CashPaymentStrategy) Pay(total float64, availableFunds float64) bool {
	return availableFunds >= total
}
