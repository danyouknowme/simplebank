package util

// Constants for all supported currencies
const (
	THB = "THB"
	USD = "USD"
	EUR = "EUR"
	WON = "WON"
	YEN = "YEN"
)

// IsSupportedCurrency returns true if the currency is supported, otherwise false
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case THB, USD, EUR, WON, YEN:
		return true
	}
	return false
}
