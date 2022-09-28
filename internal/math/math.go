package math

import "github.com/shopspring/decimal"

func CalculateOdds(odds []decimal.Decimal) decimal.Decimal {

	result := decimal.Zero

	for indx := range odds {
		result = result.Add(decimal.NewFromInt(1).Div(odds[indx]))
	}

	return result
}
