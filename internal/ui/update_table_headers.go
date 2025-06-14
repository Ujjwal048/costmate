package ui

import (
	"costmate/internal/constants"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Helper function to create a styled table cell
func createTableCell(text string, color tcell.Color, align int) *tview.TableCell {
	cell := tview.NewTableCell(text).
		SetTextColor(color).
		SetAlign(align).
		SetExpansion(1).
		SetAttributes(tcell.AttrBold).
		SetSelectable(true)
	return cell
}

// Helper function to set header cells
func SetHeaderCells(table *tview.Table, currency string) {
	headers := []struct {
		text  string
		index int
	}{
		{constants.ServiceHeader, 0},
		{fmt.Sprintf("Cost (%s)", currency), 1},
		{constants.PercentHeader, 2},
	}

	for _, header := range headers {
		table.SetCell(0, header.index, createTableCell(
			header.text,
			constants.ColumnHeaderColor,
			tview.AlignCenter,
		))
	}
}
