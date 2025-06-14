package aws

import (
	"context"
	"costmate/internal/logger"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

const (
	numPrevMonthsHistory = 3
	dateFormat           = "2006-01-02"
	monthYearFormat      = "2006-01"
)

type ServiceCost struct {
	ServiceName string
	Cost        float64
	Unit        string
	Percent     float64
}

func FetchServiceCosts(startDate, endDate time.Time) ([]ServiceCost, float64, error) {
	logger.Logger.Printf("Attempting to create AWS config...")
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, 0, fmt.Errorf("error loading AWS config: %v", err)
	}
	logger.Logger.Printf("AWS config loaded successfully.")

	svc := costexplorer.NewFromConfig(cfg)
	logger.Logger.Printf("Cost Explorer client created.")

	startDateStr := startDate.Format(dateFormat)
	endDateStr := endDate.Format(dateFormat)

	logger.Logger.Printf("Fetching costs from %s to %s (exclusive end date)...", startDateStr, endDateStr)

	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: &startDateStr,
			End:   &endDateStr,
		},
		Granularity: types.GranularityMonthly,
		Metrics:     []string{"UnblendedCost"},
		GroupBy: []types.GroupDefinition{
			{
				Type: types.GroupDefinitionTypeDimension,
				Key:  aws.String("SERVICE"),
			},
		},
	}

	result, errApi := svc.GetCostAndUsage(context.Background(), input)
	if errApi != nil {
		return nil, 0, fmt.Errorf("error fetching cost and usage: %v", errApi)
	}

	var serviceCosts []ServiceCost
	totalCost := 0.0

	if len(result.ResultsByTime) > 0 && len(result.ResultsByTime[0].Groups) > 0 {
		for _, group := range result.ResultsByTime[0].Groups {
			serviceName := "N/A"
			if len(group.Keys) > 0 {
				serviceName = group.Keys[0]
			}
			amountStr := "0.0"
			unit := "USD"
			if group.Metrics != nil {
				if metric, ok := group.Metrics["UnblendedCost"]; ok {
					amountStr = *metric.Amount
					unit = *metric.Unit
				}
			}

			costVal, _ := strconv.ParseFloat(amountStr, 64)
			totalCost += costVal

			serviceCosts = append(serviceCosts, ServiceCost{
				ServiceName: serviceName,
				Cost:        costVal,
				Unit:        unit,
			})
		}
	}

	return serviceCosts, totalCost, nil
}

// func fetchAndAppendServiceHistory(serviceName string, detailTextView *tview.TextView, app *tview.Application) {
// 	logger.Logger.Printf("Fetching history for service: %s", serviceName)

// 	cfg, err := config.LoadDefaultConfig(context.Background())
// 	if err != nil {
// 		logger.Logger.Printf("Error loading AWS config for history fetch (%s): %v", serviceName, err)
// 		app.QueueUpdateDraw(func() {
// 			currentText := detailTextView.GetText(false)
// 			newText := strings.Replace(currentText, "Fetching history for the last 3 months...\n", "", 1)
// 			detailTextView.SetText(newText + fmt.Sprintf("\n[red]Error loading AWS config for history: %v[-]", err) + "\n\n(Press Esc or b to go back)")
// 		})
// 		return
// 	}
// 	logger.Logger.Printf("AWS config loaded for history fetch: %s", serviceName)
// 	svc := costexplorer.NewFromConfig(cfg)

// 	now := time.Now()
// 	endDateForHistory := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC) // First day of current month
// 	startDateForHistory := endDateForHistory.AddDate(0, -numPrevMonthsHistory, 0)    // 3 months prior

// 	startDateStr := startDateForHistory.Format(dateFormat)
// 	endDateStr := endDateForHistory.Format(dateFormat)

// 	logger.Logger.Printf("Querying history for %s from %s to %s", serviceName, startDateStr, endDateStr)

// 	input := &costexplorer.GetCostAndUsageInput{
// 		TimePeriod: &types.DateInterval{
// 			Start: &startDateStr,
// 			End:   &endDateStr,
// 		},
// 		Granularity: types.GranularityMonthly,
// 		Metrics:     []string{"UnblendedCost"},
// 		Filter: &types.Expression{
// 			Dimensions: &types.DimensionValues{
// 				Key:    types.Dimension("SERVICE"),
// 				Values: []string{serviceName},
// 			},
// 		},
// 	}

// 	result, errApi := svc.GetCostAndUsage(context.Background(), input)
// 	app.QueueUpdateDraw(func() {
// 		currentText := detailTextView.GetText(false)
// 		// Remove the loading message more reliably
// 		lines := strings.Split(strings.TrimSpace(currentText), "\n")
// 		var newLines []string
// 		for _, line := range lines {
// 			if !strings.Contains(line, "Fetching history") {
// 				newLines = append(newLines, line)
// 			}
// 		}
// 		detailTextView.SetText(strings.Join(newLines, "\n") + "\n") // Add a newline to separate from history

// 		if errApi != nil {
// 			logger.Logger.Printf("Error fetching history for %s: %v", serviceName, errApi)
// 			fmt.Fprintf(detailTextView, "\n[red]Error fetching history: %v[-]\n", errApi)
// 		} else {
// 			logger.Logger.Printf("Successfully fetched history for %s. Results: %d", serviceName, len(result.ResultsByTime))
// 			fmt.Fprintf(detailTextView, "\nPrevious %d Months:\n", numPrevMonthsHistory)
// 			if len(result.ResultsByTime) == 0 {
// 				fmt.Fprintf(detailTextView, "No historical data found for this period.\n")
// 			} else {
// 				// Sort results by time period start date
// 				sort.Slice(result.ResultsByTime, func(i, j int) bool {
// 					return *result.ResultsByTime[i].TimePeriod.Start < *result.ResultsByTime[j].TimePeriod.Start
// 				})
// 				for _, monthlyResult := range result.ResultsByTime {
// 					periodStartStr := *monthlyResult.TimePeriod.Start
// 					periodDate, _ := time.Parse(dateFormat, periodStartStr)
// 					monthDisplay := periodDate.Format(monthYearFormat)

// 					cost := "0.00"
// 					unit := "USD"
// 					if len(monthlyResult.Total) > 0 {
// 						if metric, ok := monthlyResult.Total["UnblendedCost"]; ok {
// 							cost = *metric.Amount
// 							unit = *metric.Unit
// 						}
// 					}
// 					fmt.Fprintf(detailTextView, "%s: %s %s\n", monthDisplay, cost, unit)
// 				}
// 			}
// 		}
// 		fmt.Fprintf(detailTextView, "\n(Press Esc or b to go back)")
// 	})
// }
