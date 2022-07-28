package solution3

import (
	"fmt"
	"math"
	"strconv"
)

func Solution3(S string, B []string) []string {
	// write your code in Go 1.4
	centsInS := dollarsToCents(strToDollars(S))
	var centsInB []int

	for _, val := range B {
		centsInB = append(centsInB, dollarsToCents(strToDollars(val)))
	}

	fmt.Printf("S int cents: %v\n", centsInS)
	fmt.Printf("B int cents: %v\n", centsInB)

	resultInCents := calcDiscount(centsInS, centsInB)
	var result []string

	for _, val := range resultInCents {
		result = append(result, dollarsToStr(centsToDollars(val)))
	}

	return result
}

func calcDiscount(totalBills int, undiscountentBills []int) []int {
	return []int{} // TODO
}

func strToDollars(str string) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic("unable to parse input")
	}
	return result
}

func dollarsToCents(dollars float64) int {
	return int(math.Round(dollars * 100))
}

func centsToDollars(cents int) float64 {
	return float64(cents)
}

func dollarsToStr(dollars float64) string {
	return strconv.FormatFloat(dollars, 'g', 2, 64)
}
