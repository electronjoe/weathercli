package tests

import (
	"testing"
	"time"

	"github.com/electronjoe/weathercli/internal/api"
)

func TestClient_FetchWeatherData(t *testing.T) {
	client := api.NewClient("test-api-key")
	startDate := time.Now()

	_, err := client.FetchWeatherData(startDate)
	if err != nil {
		// Expected to fail without valid API key
		return
	}

	t.Error("Expected error with invalid API key")
}
