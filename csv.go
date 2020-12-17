package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"strconv"
)

func getAccountIndexFromNumber(accounts []Account, accountNumber string) int {
	for i, account := range accounts {
		if account.AccountNumber == accountNumber {
			return i
		}
	}
	return -1
}

func getAccountsFromCSV(filename string) (accounts []Account) {
	content, _ := ioutil.ReadFile(filename)
	reader := csv.NewReader(bytes.NewBuffer(content))
	reader.Comma = ';'
	_, err := reader.Read() // skip first line
	if err != nil {
		panic(err)
	}
	for {
		columns, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil && err != csv.ErrFieldCount {
			// "wrong number of fields" will always occur as long
			// as "lånebelopp" is not set in the konto.csv
			//panic(err)
		}

		accounts = append(accounts, Account{
			AccountNumber: columns[0],
			AccountType:   columns[1],
			TotalValue:    columns[2],
		})
	}
	return
}

func addHoldingsFromCSV(accounts []Account, filename string) []Account {
	content, _ := ioutil.ReadFile(filename)
	reader := csv.NewReader(bytes.NewBuffer(content))
	reader.Comma = ';'
	_, err := reader.Read() // skip first line
	if err != nil {
		panic(err)
	}
	for {
		columns, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		accountIndex := getAccountIndexFromNumber(accounts, columns[0])

		amount, err := strconv.ParseFloat(columns[2], 64)
		if err != nil {
			panic(err)
		}
		marketPrice, err := strconv.ParseFloat(columns[3], 64)
		if err != nil {
			panic(err)
		}
		purchasePrice, err := strconv.ParseFloat(columns[4], 64)
		if err != nil {
			panic(err)
		}

		holding := &Holding{
			Name:          columns[1],
			Amount:        amount,
			MarketPrice:   marketPrice,
			PurchasePrice: purchasePrice,
			ISIN:          columns[5],
			Currency:      columns[6],
		}
		if columns[7] == "Aktie" {
			accounts[accountIndex].Stocks = append(accounts[accountIndex].Stocks, holding)
		} else if columns[7] == "Fond" {
			accounts[accountIndex].Funds = append(accounts[accountIndex].Funds, holding)
		} else if columns[7] != "" {
			panic("invalid type " + columns[7])
		}
	}
	return accounts
}
