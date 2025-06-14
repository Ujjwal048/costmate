package utils

import (
	"costmate/internal/aws"
	"costmate/internal/logger"
	"costmate/internal/constants"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Function to fetch and display costs
func FetchCost(table *tview.Table, Month time.Time) ([]aws.ServiceCost, float64, error) {
	var err error
	var totalCost float64
	var serviceCosts []aws.ServiceCost

	if constants.UseDummyData {
		// Use dummy data
		costs := constants.ServiceCosts
		serviceCosts = make([]aws.ServiceCost, len(costs))
		for i, cost := range costs {
			serviceCosts[i] = aws.ServiceCost{
				ServiceName: cost.ServiceName,
				Cost:        cost.Cost,
				Unit:        cost.Unit,
				Percent:     cost.Percent,
			}
		}
		totalCost = constants.TotalCost
	} else {
		// Fetch real AWS costs
		startDate := time.Date(Month.Year(), Month.Month(), 1, 0, 0, 0, 0, time.UTC)
		endDate := startDate.AddDate(0, 1, 0)
		serviceCosts, totalCost, err = aws.FetchServiceCosts(startDate, endDate)
		if err != nil {
			logger.Logger.Printf("Error: %v", err)
			table.SetCell(1, 0, tview.NewTableCell(fmt.Sprintf("Error: %v", err)).
				SetTextColor(tcell.ColorRed).
				SetExpansion(3).
				SetAlign(tview.AlignCenter))
			return nil, 0, err
		}
	}
	return serviceCosts, totalCost, err
}
