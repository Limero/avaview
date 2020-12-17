package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func printHoldings(holdings []*Holding) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Värdepapper",
		"Antal",
		"+/-%",
		"Senast",
		"Inköp",
		"Markn.värde",
		"Avkastn. %",
		"Avkastn.",
	})
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
	})

	data := [][]string{}

	for _, holding := range holdings {
		totalPurchasePrice := holding.PurchasePrice * holding.Amount
		returnPercentage := ((holding.MarketPrice - totalPurchasePrice) / totalPurchasePrice) * 100

		data = append(data, []string{
			holding.Name,
			fmt.Sprintf("%.2f", holding.Amount),
			"0",
			fmt.Sprintf("%.2f", holding.MarketPrice/holding.Amount),
			fmt.Sprintf("%.0f", totalPurchasePrice),
			fmt.Sprintf("%.0f", holding.MarketPrice),
			fmt.Sprintf("%.2f", returnPercentage),
			fmt.Sprintf("%.0f", holding.MarketPrice-totalPurchasePrice),
		})
	}

	colorColumns := []int{6, 7}

	for _, v := range data {
		colors := []tablewriter.Colors{{}, {}, {}, {}, {}, {}, {}, {}}

		for _, colorColumn := range colorColumns {
			if v[colorColumn][0] == '-' {
				colors[colorColumn] = tablewriter.Colors{tablewriter.FgRedColor}
			} else {
				colors[colorColumn] = tablewriter.Colors{tablewriter.FgGreenColor}
			}
		}
		table.Rich(v, colors)
	}

	table.Render()
}
