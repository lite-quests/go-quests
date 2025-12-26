package interfaces

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
	return false
}

func (c CardPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return ""
}

type UPIPayment struct {
	VPA string
}

func (u UPIPayment) Process(amount float64) bool {
	// TODO: implement
	// Read README.md for the instructions
	return false
}

func (u UPIPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return ""
}

type CryptoPayment struct {
	Wallet  string
	Balance float64
}

func (c *CryptoPayment) Process(amount float64) bool {
	// TODO: implement
	// Read README.md for the instructions
	return false
}

func (c CryptoPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return ""
}

func Checkout(p PaymentMethod, amount float64) string {
	// TODO: implement
	// Read README.md for the instructions
	return ""
}

func DetectCrypto(p PaymentMethod) bool {
	// TODO: implement
	// Read README.md for the instructions
	return false
}
