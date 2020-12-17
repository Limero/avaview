package main

import "fmt"

func convertCurrencyToSek(fromCurrency string, fromAmount float64) float64 {
	// TODO: Make this more accurate
	switch fromCurrency {
	case "SEK":
		return fromAmount
	case "USD":
		return fromAmount * 8.30
	case "EUR":
		return fromAmount * 10.15
	}

	fmt.Println("Warning: unsupported currency", fromCurrency)
	return fromAmount
}
