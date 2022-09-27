// will break if incoem is above max float64 and maybe other edge cases
package main

import (
	"fmt"
	"math"
)

func main() {
	// income := 20000.
	maxIncome := 10000000.
	incremIncome := 15000.
	//tax bracket from this source: https://taxfoundation.org/2022-tax-brackets/
	brackets := []struct {
		min     float64
		max     float64
		percent float64
	}{
		{0, 10275, .10}, {10275, 41775, .12}, {41775, 89075, .22}, {89075, 170050, .24}, {170050, 215950, .32}, {215950, 539900, .35}, {539900, math.MaxFloat64, .37},
	}
	type absoluteTax struct {
		income float64
		tax    float64
	}
	absTaxes := []absoluteTax{}
	for income := 1.; income < maxIncome; income += incremIncome {
		tax := 0.
		calculated := 0.
		for _, bracket := range brackets {
			min := bracket.min
			max := bracket.max
			percent := bracket.percent
			inBracket := income - min
			if inBracket > max-min {
				inBracket = max - min
			}
			tax += inBracket * percent
			calculated += inBracket
			remainder := income - calculated
			fmt.Printf("Bracket tax: %v, income above current bracket: %v \n", inBracket*percent, remainder)
			if remainder < 0.0001 {
				break
			}
		}
		absTaxes = append(absTaxes, absoluteTax{income, tax})
		fmt.Printf("total taxed: %v  percent of income: %v \n", tax, (tax/income)*100)
	}
	// fmt.Println(absTaxes)
	for _, absTax := range absTaxes {
		for tax := absTax.tax; tax > 0; tax = tax - 100000 {
			fmt.Printf("|")
		}
		fmt.Println("")

	}
}
