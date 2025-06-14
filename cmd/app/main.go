package main

import (
	"costmate/internal/bootstrap"
	"costmate/internal/constants"
	"costmate/internal/handler"
	"costmate/internal/logger"
)

func main() {

	bootstrap.InitDependencies()
	currentMonth := constants.CurrentMonth
	currency := constants.DefaultCurrency
	app, flex, table, info := bootstrap.LoadInitialView()

	serviceCosts, totalCost, err := bootstrap.GetInitialCost(table, currentMonth, currency)
	if err != nil {
		logger.Error("Failed to fetch initial costs", err)
		return
	}
	// Set initial selection
	table.Select(1, 0)

	// Setup keyboard handlers
	handler.SetupKeyboardHandlers(app, flex, table, info, &serviceCosts, &totalCost, &currency, &currentMonth)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		logger.Fatal("Application error", err)
	}
}
