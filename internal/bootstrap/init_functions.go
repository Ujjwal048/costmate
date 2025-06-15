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
		logger.Error("Failed to fetch dollar rate", err)
	}
}

func GetInitialCost(table *tview.Table, currentMonth time.Time, currency string) ([]aws.ServiceCost, float64, error) {
	var err error
	serviceCosts, totalCost, err := utils.FetchCost(currentMonth)
	if err != nil {
		logger.Error("Failed to fetch initial costs", err)
		return nil, 0, err
	}

	utils.UpdateTableWithCosts(table, serviceCosts, totalCost, currency, currentMonth)
	return serviceCosts, totalCost, nil
}
