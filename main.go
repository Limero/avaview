package main

import (
	"fmt"
)

func main() {
	accounts := getAccountsFromCSV("konto.csv")
	accounts = addHoldingsFromCSV(accounts, "positioner.csv")

	for _, account := range accounts {
		fmt.Println(account.AccountType)

		if len(account.Stocks) > 0 {
			fmt.Println("Aktier")

			for _, stock := range account.Stocks {
				stockId := getStockIdFromName(stock.Name)
				updatedStock := getStockFromId(stockId)
				stock.MarketPrice = convertCurrencyToSek(stock.Currency, updatedStock.LastPrice*stock.Amount)
			}

			printHoldings(account.Stocks)
		}

		if len(account.Funds) > 0 {
			fmt.Println("Fonder")

			for _, fund := range account.Funds {
				fundId := getFundIdFromName(fund.Name)
				updatedFund := getFundFromId(fundId)
				fund.MarketPrice = updatedFund.Nav * fund.Amount
			}

			printHoldings(account.Funds)
		}

		fmt.Println("")
	}
}
