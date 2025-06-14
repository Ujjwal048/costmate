package ui

import (
	"fmt"

	"costmate/internal/aws"
	"costmate/internal/constants"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetServiceCostCells(table *tview.Table, row int, cost aws.ServiceCost, displayCost float64) {
	cells := []struct {
		text  string
		index int
	}{
		{cost.ServiceName, 0},
		{fmt.Sprintf("%.2f", displayCost), 1},
		{fmt.Sprintf("%.2f%%", cost.Percent), 2},
	}

	for _, cell := range cells {
		table.SetCell(row, cell.index, createTableCell(
			cell.text,
			constants.ContentColor,
			tview.AlignCenter,
		))
	}
}

// Helper function to set total row cells
func SetTotalRowCells(table *tview.Table, row int, totalCost float64) {
	cells := []struct {
		text  string
		index int
	}{
		{"TOTAL", 0},
		{fmt.Sprintf("%.2f", totalCost), 1},
		{"100.00%", 2},
	}

	for _, cell := range cells {
		table.SetCell(row, cell.index, createTableCell(
			cell.text,
			tcell.ColorRed,
			tview.AlignCenter,
		))
	}
}
