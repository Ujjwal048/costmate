package handler

import (
	"costmate/internal/aws"
	"costmate/internal/constants"
	"costmate/internal/logger"
	"costmate/internal/modals"
	"costmate/internal/utils"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetupKeyboardHandlers(app *tview.Application, flex *tview.Flex, table *tview.Table, info *tview.TextView, serviceCosts *[]aws.ServiceCost, totalCost *float64, currency *string, currentMonth *time.Time) {
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			row, _ := table.GetSelection()
			if row > 1 {
				table.Select(row-1, 0)
			}
			return nil
		case tcell.KeyDown:
			row, _ := table.GetSelection()
			if row < len(*serviceCosts)+1 {
				table.Select(row+1, 0)
			}
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			case 'p':
				modals.SwitchProfile(app, flex, table, info, func(profileName string) {
					var err error
					*serviceCosts, *totalCost, err = utils.FetchCost(*currentMonth)
					if err != nil {
						logger.Error("Error fetching costs: %v", err)
						return
					}
					utils.UpdateTableWithCosts(table, *serviceCosts, *totalCost, constants.DefaultCurrency, *currentMonth)
				})
				return nil

			case 'c':
				if *currency == "USD" {
					*currency = "INR"
				} else {
					*currency = "USD"
				}
				logger.Info("Switched currency to: %s", *currency)
				utils.UpdateTableWithCosts(table, *serviceCosts, *totalCost, *currency, *currentMonth)

			case 's':
				utils.SortServiceCosts(table, *serviceCosts, *totalCost, *currency, *currentMonth)

			case 'm':
				handleMonthSwitch := func(selectedMonth time.Time) {
					var err error
					*serviceCosts, *totalCost, err = utils.FetchCost(selectedMonth)
					if err != nil {
						logger.Error("Error fetching costs: %v", err)
						return
					}
					*currentMonth = selectedMonth
					utils.UpdateTableWithCosts(table, *serviceCosts, *totalCost, constants.DefaultCurrency, selectedMonth)
				}
				modals.SwitchMonth(app, flex, table, info, handleMonthSwitch)
				return nil
			}
			return nil
		}
		return event
	})
}
