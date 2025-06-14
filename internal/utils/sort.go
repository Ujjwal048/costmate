package utils

import (
	"costmate/internal/aws"
	"costmate/internal/logger"
	
	"sort"
	"time"

	"github.com/rivo/tview"
)

func SortServiceCosts(table *tview.Table, serviceCosts []aws.ServiceCost, totalCost float64, currency string, currentMonth time.Time ) {

	logger.Logger.Printf("Sorting services by cost")
	// Sort services by cost in descending order
	sort.Slice(serviceCosts, func(i, j int) bool {
		return serviceCosts[i].Cost > serviceCosts[j].Cost
	})

	// Recalculate percentages after sorting
	for i := range serviceCosts {
		serviceCosts[i].Percent = (serviceCosts[i].Cost / totalCost) * 100
	}
	UpdateTableWithCosts(table, serviceCosts, totalCost, currency, currentMonth)

}
