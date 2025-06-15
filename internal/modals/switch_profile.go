package modals

import (
	"costmate/internal/aws"
	"costmate/internal/logger"
	"costmate/internal/ui"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Function to show profile selection
func SwitchProfile(app *tview.Application, flex *tview.Flex, table *tview.Table, info *tview.TextView, onSelect func(string)) {

	// Function to show error modal
	showErrorModal := func(app *tview.Application, flex *tview.Flex, message string) {
		logger.Error("Error", fmt.Errorf(message))
		modal := tview.NewModal().
			SetText(message).
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.SetRoot(flex, true)
			})
		app.SetRoot(modal, true)
	}

	profiles, err := aws.GetAvailableProfiles()
	if err != nil {
		showErrorModal(app, flex, fmt.Sprintf("Error loading profiles: %v", err))
		return
	}

	// Function to format profile list with proper alignment
	formatProfileList := func(selectedIndex int) string {
		var text strings.Builder
		text.WriteString("[::b]Select AWS Profile[::-]\n\n")
		for i, profile := range profiles {
			if i == selectedIndex {
				text.WriteString(fmt.Sprintf("  [::r]%-15s[::-]\n", profile.Name))
			} else {
				text.WriteString(fmt.Sprintf("  %-15s\n", profile.Name))
			}
		}
		return text.String()
	}

	// Function to switch profile
	switchProfile := func(profileName string) {
		if err := aws.SwitchProfile(profileName); err != nil {
			showErrorModal(app, flex, fmt.Sprintf("Error switching profile: %v", err))
			return
		}

		logger.Info("Switched to profile: %s", profileName)
		info = ui.UpdateInfo(info, profileName, app, flex)
		onSelect(profileName)
	}

	currentSelection := 0
	modal := tview.NewModal().
		SetText(formatProfileList(currentSelection))

	modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if currentSelection > 0 {
				currentSelection--
				modal.SetText(formatProfileList(currentSelection))
			}
			return nil
		case tcell.KeyDown:
			if currentSelection < len(profiles)-1 {
				currentSelection++
				modal.SetText(formatProfileList(currentSelection))
			}
			return nil
		case tcell.KeyEnter:
			if currentSelection < len(profiles) {
				switchProfile(profiles[currentSelection].Name)
			}
			return nil
		case tcell.KeyEscape:
			app.SetRoot(flex, true)
			return nil
		}
		return event
	})

	app.SetRoot(modal, true)
}
