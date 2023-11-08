package patterns

type CardPaymentStrategy struct{}

func (cs CardPaymentStrategy) Pay(total float64, availableFunds float64) bool {
	return availableFunds >= total
}
