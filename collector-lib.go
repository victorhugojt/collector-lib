package main

import (
	"fmt"
	"strconv"
)

func main() {
	const percentageOfRevenue = 0.15
	message := "Victor you have : "
	numberOfRecords := 1200
	var stimatedTotalValue float64 = 60000 + (60000 * percentageOfRevenue)

	fmt.Println(message + strconv.Itoa(numberOfRecords) + " records with a stimated value of " + strconv.FormatFloat(stimatedTotalValue, 'f', -1, 64))
}
