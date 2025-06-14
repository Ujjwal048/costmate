package bootstrap

import (
	"costmate/internal/constants"
	"costmate/internal/logger"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func LoadInitialView() (*tview.Application, *tview.Flex, *tview.Table, *tview.TextView) {

	// currentMonth = time.Now()
	app := tview.NewApplication()
	flex := tview.NewFlex()

	logo := tview.NewTextView().
		SetText(constants.Logo).
		SetTextColor(constants.LogoColor).
		SetTextAlign(tview.AlignRight)

	info := tview.NewTextView().
		SetText(constants.InfoText).
		SetTextColor(constants.InfoTextColor).
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetWrap(true)

	helpbox := tview.NewTextView().
		SetText(constants.HelpText).
		SetTextColor(tcell.ColorBlueViolet).
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetWrap(true)

	logoFrame := tview.NewFrame(logo)
	infoFrame := tview.NewFrame(info)
	helpFrame := tview.NewFrame(helpbox)

	topFlex := tview.NewFlex().
		AddItem(infoFrame, 0, 1, false).
		AddItem(helpFrame, 0, 1, false).
		AddItem(logoFrame, 0, 1, false)

	table := tview.NewTable().SetBorders(true).SetFixed(1, 1)
	table.SetBorder(true).
		SetBorderColor(tcell.ColorTeal).
		SetTitleColor(constants.InfoTextColor).
		SetBorderPadding(0, 0, 0, 0)

	table.SetSelectable(true, false).SetBackgroundColor(tcell.ColorDefault)

	table.SetCell(0, 0, tview.NewTableCell(constants.ServiceHeader).
		SetTextColor(constants.ColumnHeaderColor).
		SetAttributes(tcell.AttrBold).
		SetAlign(tview.AlignCenter).
		SetExpansion(1))
	table.SetCell(0, 1, tview.NewTableCell(constants.CostHeader).
		SetTextColor(constants.ColumnHeaderColor).
		SetAttributes(tcell.AttrBold).
		SetAlign(tview.AlignCenter).
		SetExpansion(1))
	table.SetCell(0, 2, tview.NewTableCell(constants.PercentHeader).
		SetTextColor(constants.ColumnHeaderColor).
		SetAttributes(tcell.AttrBold).
		SetAlign(tview.AlignCenter).
		SetExpansion(1))

	innerFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	innerFlex.AddItem(topFlex, 9, 1, false)
	innerFlex.AddItem(table, 0, 3, true)

	flex.AddItem(innerFlex, 0, 2, true).
		SetBorderPadding(0, 0, 0, 0)

	// Set focus to table
	app.SetFocus(table)
	logger.Info("Application UI initialized")

	return app, flex, table, info

}
