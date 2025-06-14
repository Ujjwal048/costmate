package utils

import (
	"costmate/internal/logger"
	"encoding/json"
	"io"
	"net/http"
)

func ConvertDollarToRupee(value float64) float64 {
	rate := GetCurrentDollarRate()
	return value * rate
}

func GetCurrentDollarRate() float64 {
	resp, err := http.Get("https://api.frankfurter.app/latest?from=USD&to=INR")
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0
	}
	value := data["rates"].(map[string]interface{})["INR"].(float64)
	logger.Logger.Printf("Dollar Rate: %f", value)
	return value
}
