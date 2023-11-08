package patterns

// Funds represents the available funds for the user
type Funds struct {
	Amount float64
}

var fundsInstance *Funds

// GetFundsInstance returns the singleton instance of available funds
func GetFundsInstance() *Funds {
	if fundsInstance == nil {
		fundsInstance = &Funds{Amount: 20.0} // Initialize with $20
	}
	return fundsInstance
}
