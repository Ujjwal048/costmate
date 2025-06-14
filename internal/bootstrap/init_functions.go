package bootstrap

import (
	"costmate/internal/aws"
	"costmate/internal/logger"
	"costmate/internal/utils"
	"log"
	"time"

	"github.com/rivo/tview"
)

func InitDependencies() {
	if err := logger.Initialize(); err != nil {
		log.Fatal("Error initializing logger:", err)
	}
	defer logger.Close()

	// Initialize dollar rate
	if err := utils.GetDollarRate(); err != nil {
		logger.Logger.Printf("Warning: Failed to fetch dollar rate: %v", err)
	}
}

func GetInitialCost(table *tview.Table, currentMonth time.Time, currency string) ([]aws.ServiceCost, float64, error) {
	var err error
	serviceCosts, totalCost, err := utils.FetchCost(currentMonth)
	if err != nil {
		logger.Logger.Printf("Error Fetching Initial Cost: %v", err)
		return nil, 0, err
	}

	utils.UpdateTableWithCosts(table, serviceCosts, totalCost, currency, currentMonth)
	return serviceCosts, totalCost, nil
}
