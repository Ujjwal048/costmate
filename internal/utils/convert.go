package utils

import (
	"costmate/internal/logger"
	"encoding/json"
	"io"
	"net/http"
)

var currentRate float64

// InitializeDollarRate fetches and stores the current dollar rate
func GetDollarRate() error {
	resp, err := http.Get("https://api.frankfurter.app/latest?from=USD&to=INR")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	currentRate = data["rates"].(map[string]interface{})["INR"].(float64)
	logger.Info("Dollar Rate: %f", currentRate)
	return nil
}

func ConvertDollarToRupee(value float64) float64 {
	return value * currentRate
}
