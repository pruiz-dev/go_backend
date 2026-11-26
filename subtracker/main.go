package main

import "fmt"

func main() {
	serviceName := "Test service"
	cost := 299.5
	monthsCount := 12
	isActive := true
	totalCost := calcTotalCost(cost, monthsCount)

	const rowHeaderFmt = "%-12s  %-13s  %-12s  %-22s  %-23s\n"
	const rowValuesFmt = "%-12s  %-13.2f  %-12d  %-22t  %-23.2f\n\n"
	fmt.Printf(rowHeaderFmt, "Service name", "Cost by month", "Months count", "Subscription is active", "Total cost (all months)")
	fmt.Printf(rowHeaderFmt, "------------", "-------------", "------------", "----------------------", "-----------------------")
	fmt.Printf(rowValuesFmt, serviceName, cost, monthsCount, isActive, totalCost)
}

func calcTotalCost(cost float64, monthsCount int) float64 {
	return float64(monthsCount) * cost
}
