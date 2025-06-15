package constants

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
)

// HelpText contains the help information displayed in the UI
const HelpText = `  [::b] <p>   Switch Profile
   <c>   Convert to INR/USD
   <s>   Sort by Cost
   <m>   Filter by Month [::-]`

const InfoText = `  [::b] AWS Profile: default [::-]`

// Colors used in the UI
const (
	// Text colors
	InfoTextColor     = tcell.ColorYellow
	ColumnHeaderColor = tcell.ColorYellow
	ContentColor      = tcell.ColorLightSkyBlue
	TotalColor        = tcell.ColorRed
	HelpTextColor     = tcell.ColorBlueViolet
	LogoColor         = tcell.ColorTeal
	ModalColor        = tcell.ColorLightBlue
)

// Table headers
const (
	ServiceHeader = "Service"
	CostHeader    = "Cost (USD)"
	PercentHeader = "Percent %"
)

// Logo is loaded from the assets file
var Logo string

func init() {
	// Read logo from file
	logoPath := filepath.Join("assets", "logo.txt")
	logoBytes, err := os.ReadFile(logoPath)
	if err != nil {
		// If file can't be read, use a simple fallback logo
		Logo = "CostMate"
	} else {
		Logo = string(logoBytes)
	}
}
