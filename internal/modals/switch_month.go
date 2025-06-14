package modals

import (
	"costmate/internal/logger"
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SwitchMonth(app *tview.Application, flex *tview.Flex, table *tview.Table, info *tview.TextView, onSelect func(time.Time)) {
	// Function to show month selection
	logger.Info("Opening month selection")
	months := make([]time.Time, 12)
	now := time.Now()
	// Get the first day of the current month
	currentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 12; i++ {
		months[i] = currentMonth.AddDate(0, -i, 0)
	}

	// Function to format month list with proper alignment
	formatMonthList := func(selectedIndex int) string {
		var text strings.Builder
		text.WriteString("[::b]Select Month[::-]\n\n")
		for i, month := range months {
			if i == selectedIndex {
				text.WriteString(fmt.Sprintf("  [::r]%-15s[::-]\n", month.Format("January 2006")))
			} else {
				text.WriteString(fmt.Sprintf("  %-15s\n", month.Format("January 2006")))
			}
		}
		return text.String()
	}

	// Track current selection
	currentSelection := 0
	totalItems := len(months)

	// Set initial formatted text
	modal := tview.NewModal().
		SetText(formatMonthList(currentSelection))

	// Add keyboard navigation
	modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if currentSelection > 0 {
				currentSelection--
				modal.SetText(formatMonthList(currentSelection))
			}
			return nil
		case tcell.KeyDown:
			if currentSelection < totalItems-1 {
				currentSelection++
				modal.SetText(formatMonthList(currentSelection))
			}
			return nil
		case tcell.KeyEnter:
			onSelect(months[currentSelection])
			app.SetRoot(flex, true)
			return nil
		case tcell.KeyEscape:
			app.SetRoot(flex, true)
			return nil
		}
		return event
	})

	app.SetRoot(modal, true)
}
