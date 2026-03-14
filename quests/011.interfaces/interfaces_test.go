package interfaces

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the interfaces Quest 🎉", colorReset)
	}
	os.Exit(code)
}

/*
========================
Interface Satisfaction
========================
*/

func TestPaymentMethodInterface(t *testing.T) {
	var _ PaymentMethod = (*CardPayment)(nil)
	var _ PaymentMethod = UPIPayment{}
	var _ PaymentMethod = (*CryptoPayment)(nil)
}

/*
========================
CardPayment.Process
========================*/

func TestCardPaymentProcess(t *testing.T) {
	tests := []struct {
		startLimit float64
		amount     float64
		wantOK     bool
		wantLimit  float64
	}{
		{100, 40, true, 60},
		{50, 60, false, 50},
		{0, 0, true, 0},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			c := &CardPayment{Limit: tt.startLimit}
			ok := c.Process(tt.amount)

			if ok != tt.wantOK {
				t.Fatalf("Process() = %v, want %v", ok, tt.wantOK)
			}

			if c.Limit != tt.wantLimit {
				t.Fatalf("Limit = %v, want %v", c.Limit, tt.wantLimit)
			}
		})
	}
}

/*
========================
UPIPayment.Process
========================
*/

func TestUPIPaymentProcess(t *testing.T) {
	tests := []struct {
		amount float64
	}{
		{0},
		{100},
		{9999},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			u := UPIPayment{VPA: "test@upi"}
			if ok := u.Process(tt.amount); !ok {
				t.Fatalf("Process() = false, want true")
			}
		})
	}
}

/*
========================
CryptoPayment.Process
========================
*/

func TestCryptoPaymentProcess(t *testing.T) {
	tests := []struct {
		startBal float64
		amount   float64
		wantOK   bool
		wantBal  float64
	}{
		{10, 3, true, 7},
		{5, 10, false, 5},
		{1, 1, true, 0},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			c := &CryptoPayment{Balance: tt.startBal}
			ok := c.Process(tt.amount)

			if ok != tt.wantOK {
				t.Fatalf("Process() = %v, want %v", ok, tt.wantOK)
			}

			if c.Balance != tt.wantBal {
				t.Fatalf("Balance = %v, want %v", c.Balance, tt.wantBal)
			}
		})
	}
}

/*
========================
Provider
========================
*/

func TestProvider(t *testing.T) {
	tests := []struct {
		name     string
		payment  PaymentMethod
		wantProv string
	}{
		{"card", &CardPayment{}, "CARD"},
		{"upi", UPIPayment{}, "UPI"},
		{"crypto", &CryptoPayment{}, "CRYPTO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.payment.Provider(); got != tt.wantProv {
				t.Fatalf("Provider() = %q, want %q", got, tt.wantProv)
			}
		})
	}
}

/*
========================
Checkout
========================
*/

func TestCheckout(t *testing.T) {
	tests := []struct {
		name    string
		payment PaymentMethod
		amount  float64
		wantOK  bool
	}{
		{"card-success", &CardPayment{Limit: 100}, 50, true},
		{"card-fail", &CardPayment{Limit: 10}, 50, false},
		{"upi", UPIPayment{VPA: "a@upi"}, 500, true},
		{"crypto", &CryptoPayment{Balance: 5}, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := Checkout(tt.payment, tt.amount)

			if tt.wantOK && !strings.Contains(msg, "successful") {
				t.Fatalf("Checkout() = %q, expected success message", msg)
			}

			if !tt.wantOK && !strings.Contains(msg, "failed") {
				t.Fatalf("Checkout() = %q, expected failure message", msg)
			}
		})
	}
}

/*
========================
DetectCrypto
========================
*/

func TestDetectCrypto(t *testing.T) {
	tests := []struct {
		name string
		p    PaymentMethod
		want bool
	}{
		{"crypto", &CryptoPayment{}, true},
		{"card", &CardPayment{}, false},
		{"upi", UPIPayment{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCrypto(tt.p); got != tt.want {
				t.Fatalf("DetectCrypto() = %v, want %v", got, tt.want)
			}
		})
	}
}
