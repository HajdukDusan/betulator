package main

import (
	"betulator/pkg/crawlers/meridianbet"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shopspring/decimal"
)

func main() {

	bets, err := meridianbet.GetBets()
	if err != nil {
		fmt.Println(err)
	}

	for key, val := range bets {
		printBet(key[:], val[:])
	}

}

func printBet(winners []string, odds []decimal.Decimal) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Outcome", "Odds"})

	winnersRow := make([]interface{}, len(winners))
	for i := range winners {
		winnersRow[i] = winners[i]
	}
	t.AppendHeader(winnersRow)

	oddsRow := make([]interface{}, len(odds))
	for i := range odds {
		oddsRow[i] = odds[i]
	}
	t.AppendRow(oddsRow)
	t.AppendSeparator()
	t.AppendFooter(table.Row{"", "TOTAL ODDS", 10000})

	t.SetStyle(table.StyleLight)
	t.Render()
}
