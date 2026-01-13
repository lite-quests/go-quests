package solutions

type PaymentMethod interface {
	Process(amount float64) bool
	Provider() string
}

type CardPayment struct {
	CardNumber string
	Limit      float64
}

func (c *CardPayment) Process(amount float64) bool {
	// TODO: implement
	// Read README.md for the instructions
	if amount > c.Limit {
		return false
	}
	c.Limit = c.Limit - amount
	return true
}

func (c CardPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return "CARD"
}

type UPIPayment struct {
	VPA string
}

func (u UPIPayment) Process(amount float64) bool {
	// TODO: implement
	// Read README.md for the instructions
	return true
}

func (u UPIPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return "UPI"
}

type CryptoPayment struct {
	Wallet  string
	Balance float64
}

func (c *CryptoPayment) Process(amount float64) bool {
	// TODO: implement
	// Read README.md for the instructions
	if amount > c.Balance {
		return false
	}
	c.Balance = c.Balance - amount
	return true
}

func (c CryptoPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return "CRYPTO"
}

func Checkout(p PaymentMethod, amount float64) string {
	// TODO: implement
	// Read README.md for the instructions
	if p.Process(amount) {
		return "Payment successful via " + p.Provider()
	}
	return "Payment failed via " + p.Provider()

}

func DetectCrypto(p PaymentMethod) bool {
	// TODO: implement
	// Read README.md for the instructions

	if p.Provider() == "CRYPTO" {
		return true
	}
	return false
}
