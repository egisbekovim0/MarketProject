package patterns

// Funds represents the available funds for the user
type Funds struct {
	Amount float64
}

type FundsCard struct {
	Amount float64
}

var fundsInstance *Funds

func GetFundsInstance() *Funds {
	if fundsInstance == nil {
		fundsInstance = &Funds{Amount: 20.0} // Initialize with $20
	}
	return fundsInstance
}

var fundsCardInstance *FundsCard

func GetFundsCardInstance() *FundsCard {
	if fundsCardInstance == nil {
		fundsCardInstance = &FundsCard{Amount: 100.0}
	}
	return fundsCardInstance
}
