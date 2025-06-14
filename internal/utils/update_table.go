package utils

import (
	"fmt"
	"time"

	"costmate/internal/aws"
	"costmate/internal/ui"

	"github.com/rivo/tview"
)

func UpdateTableWithCosts(table *tview.Table, serviceCosts []aws.ServiceCost, totalCost float64, currency string, currentMonth time.Time) {
	// Clear the table first
	table.Clear()
	ui.SetHeaderCells(table, currency)

	// Set up conversion function based on currency
	var convertCost func(float64) float64
	if currency == "INR" {
		convertCost = ConvertDollarToRupee
		totalCost = convertCost(totalCost)
	} else {
		convertCost = func(cost float64) float64 { return cost }
	}

	// Display service costs
	for i, cost := range serviceCosts {
		displayCost := convertCost(cost.Cost)
		cost.Percent = (displayCost / totalCost) * 100
		ui.SetServiceCostCells(table, i+1, cost, displayCost)
	}

	// Add total row
	ui.SetTotalRowCells(table, len(serviceCosts)+1, totalCost)

	// Set table title
	title := fmt.Sprintf("AWS Costs (%s)", currentMonth.Format("January 2006"))
	table.SetTitle(title)
}
