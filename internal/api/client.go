package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
    apiKey     string
    httpClient *http.Client
}

type WeatherData struct {
    Date         time.Time
    TempMax      float64
    FeelsLikeMax float64
    TempMin      float64
    FeelsLikeMin float64
    Precip       float64
    PrecipType   string
    WindGust     float64
    WindSpeed    float64
    CloudCover   float64
    Conditions   string
}

func NewClient(apiKey string) *Client {
    return &Client{
        apiKey: apiKey,
        httpClient: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (c *Client) FetchWeatherData(startDate time.Time) ([]WeatherData, error) {
    url := fmt.Sprintf(
        "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%f,%f/%s/%s?key=%s",
        40.7128, // Default to NYC coordinates
        -74.0060,
        startDate.Format("2006-01-02"),
        startDate.AddDate(0, 0, 6).Format("2006-01-02"),
        c.apiKey,
    )

    resp, err := c.httpClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch weather data: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
    }

    var result []WeatherData
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode API response: %v", err)
    }

    return result, nil
}
