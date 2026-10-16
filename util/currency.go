package util

var supportedCurrencies = map[string]string{
	"USD": "USD",
	"EUR": "EUR",
	"GBP": "GBP",
	"JPY": "JPY",
	"NZD": "NZD",
	"CAD": "CAD",
	"CHF": "CHF",
	"CHW": "CHW",
	"CNY": "CNY",
	"HKD": "HKD",
	"RUB": "RUB",
	"KES": "KES",
}

func IsSupportedCurrency(currency string) bool {
	_, ok := supportedCurrencies[currency]
	return ok
}
