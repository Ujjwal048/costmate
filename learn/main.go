package main

import (
	"fmt"
	"time"
)

func main() {
	currentMonth := time.Now()
	startDate := time.Date(currentMonth.Year(), currentMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)
	fmt.Println(startDate)
	fmt.Println(endDate)
	handleMonthSwitch := func(selectedMonth time.Time) {
		currentMonth = selectedMonth
		fmt.Println(currentMonth)
	}
	handleMonthSwitch(time.Now())
}
