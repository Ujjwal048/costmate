package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func UpdateInfo(info *tview.TextView, profileName string, app *tview.Application, flex *tview.Flex) *tview.TextView {
	info.SetText(fmt.Sprintf("  [::b]AWS Profile: %s[::-]", profileName))
	app.SetRoot(flex, true)
	return info
}
