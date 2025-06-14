package main

import (
	"log"
	"time"

	"costmate/internal/bootstrap"
	"costmate/internal/constants"
	"costmate/internal/logger"
	"costmate/internal/modals"
	"costmate/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// Initialize logger
	if err := logger.Initialize(); err != nil {
		log.Fatal("Error initializing logger:", err)
	}
	defer logger.Close()

	currentMonth := constants.CurrentMonth
	currency := constants.DefaultCurrency
	app, flex, table, info := bootstrap.LoadInitialView()

	var err error
	serviceCosts, totalCost, err := utils.FetchCost(table, currentMonth)
	if err != nil {
		logger.Logger.Printf("Error: %v", err)
		return
	}

	utils.UpdateTableWithCosts(table, serviceCosts, totalCost, currency, currentMonth)

	// Set initial selection
	selectedRow := 1
	table.Select(selectedRow, 0)

	// Handle keyboard events
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if selectedRow > 1 { // Don't go above the first data row
				selectedRow--
				table.Select(selectedRow, 0)
			}
			return nil
		case tcell.KeyDown:
			if selectedRow < len(serviceCosts)+1 { // Allow navigation to totals row
				selectedRow++
				table.Select(selectedRow, 0)
			}
			return nil
		case tcell.KeyRune:
			switch event.Rune() {

			case 'p':
				modals.SwitchProfile(app, flex, table, info, func(profileName string) {
					var err error
					serviceCosts, totalCost, err = utils.FetchCost(table, currentMonth)
					if err != nil {
						logger.Logger.Printf("Error fetching costs: %v", err)
						return
					}
					utils.UpdateTableWithCosts(table, serviceCosts, totalCost, constants.DefaultCurrency, currentMonth)
				})
				return nil

			case 'c':
				// Toggle currency
				if currency == "USD" {
					currency = "INR"
				} else {
					currency = "USD"
				}
				logger.Logger.Printf("Currency switched currency to: %s", currency)
				utils.UpdateTableWithCosts(table, serviceCosts, totalCost, currency, currentMonth)

			case 's':
				utils.SortServiceCosts(table, serviceCosts, totalCost, currency, currentMonth)

			case 'm':
				// Handle month switch
				handleMonthSwitch := func(selectedMonth time.Time) {
					var err error
					serviceCosts, totalCost, err = utils.FetchCost(table, selectedMonth)
					if err != nil {
						logger.Logger.Printf("Error fetching costs: %v", err)
						return
					}
					utils.UpdateTableWithCosts(table, serviceCosts, totalCost, constants.DefaultCurrency, selectedMonth)
				}
				modals.SwitchMonth(app, flex, table, info, handleMonthSwitch)
				return nil
			}
			return nil
		}
		return event
	})

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		logger.Logger.Printf("Fatal error: %v", err)
		log.Fatal(err)
	}
}
