package main

import (
	"fmt"
	"math"
)

func currencyCalculate(oneDollarCoin, fiftyCentCoin, twentyCentCoin, tenCentCoin, fiveCentCoin float64) (float64, float64, float64) {
	oneDollar := 1 * oneDollarCoin
	fiftyCent := 0.5 * fiftyCentCoin
	twentyCent := 0.2 * twentyCentCoin
	tenCent := 0.1 * tenCentCoin
	fiveCent := 0.05 * fiveCentCoin
	totalAmount := oneDollar + fiftyCent + twentyCent + tenCent + fiveCent
	twoDollarNotes := totalAmount / 2
	remainder := totalAmount - (math.Floor(twoDollarNotes) * 2)
	return totalAmount, twoDollarNotes, remainder
}

func main() {
	var oneDollarCoin float64
	var fiftyCentCoin float64
	var twentyCentCoin float64
	var tenCentCoin float64
	var fiveCentCoin float64

	fmt.Println("How many one dollar coins? ")
	fmt.Scanln(&oneDollarCoin)
	fmt.Println("How many 50 cent coins? ")
	fmt.Scanln(&fiftyCentCoin)
	fmt.Println("How many 20 cent coins? ")
	fmt.Scanln(&twentyCentCoin)
	fmt.Println("How many 10 cent coins? ")
	fmt.Scanln(&tenCentCoin)
	fmt.Println("How many 5 cent coins? ")
	fmt.Scanln(&fiveCentCoin)
	totalAmount, twoDollarNotes, remainder := currencyCalculate(oneDollarCoin, fiftyCentCoin, twentyCentCoin, tenCentCoin, fiveCentCoin)
	fmt.Println("Total amount is $", totalAmount, "\nNumber of two dollar notes are", math.Floor(twoDollarNotes), "\nRemaining money is $", math.Round(remainder*100)/100)

}
