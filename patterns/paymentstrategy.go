package patterns

type PaymentStrategy interface {
	Pay(total float64, availableFunds float64) bool
}
