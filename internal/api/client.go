package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

// This matches the actual VisualCrossing API response structure
type APIResponse struct {
	QueryCost       int       `json:"queryCost"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	ResolvedAddress string    `json:"resolvedAddress"`
	Address         string    `json:"address"`
	Timezone        string    `json:"timezone"`
	Days            []DayData `json:"days"`
}

type DayData struct {
	Datetime     string   `json:"datetime"`
	TempMax      float64  `json:"tempmax"`
	FeelsLikeMax float64  `json:"feelslikemax"`
	TempMin      float64  `json:"tempmin"`
	FeelsLikeMin float64  `json:"feelslikemin"`
	Precip       float64  `json:"precip"`
	PrecipType   []string `json:"preciptype"`
	WindGust     float64  `json:"windgust"`
	WindSpeed    float64  `json:"windspeed"`
	CloudCover   float64  `json:"cloudcover"`
	Conditions   string   `json:"conditions"`
}

// WeatherData is our internal representation
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

func init() {
	// Set up logging
	log.SetLevel(log.WarnLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
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
	// Using Nashville, TN coordinates as default (modify as needed)
	url := fmt.Sprintf(
		"https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/Nashville/%s/%s?key=%s&include=days",
		startDate.Format("2006-01-02"),
		startDate.AddDate(0, 0, 6).Format("2006-01-02"),
		c.apiKey,
	)

	log.Debugf("Requesting URL: %s", url)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Errorf("API request failed with status: %d, body: %s", resp.StatusCode, string(bodyBytes))
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Log the raw response
	log.Debugf("Raw API Response: %s", string(bodyBytes))

	var apiResp APIResponse
	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %v", err)
	}

	// Convert API response to our internal format
	var result []WeatherData
	for _, day := range apiResp.Days {
		date, err := time.Parse("2006-01-02", day.Datetime)
		if err != nil {
			log.Warnf("Failed to parse date %s: %v", day.Datetime, err)
			continue
		}

		precipType := "None"
		if day.PrecipType != nil && len(day.PrecipType) > 0 {
			precipType = strings.Join(day.PrecipType, ",")
		}

		result = append(result, WeatherData{
			Date:         date,
			TempMax:      day.TempMax,
			FeelsLikeMax: day.FeelsLikeMax,
			TempMin:      day.TempMin,
			FeelsLikeMin: day.FeelsLikeMin,
			Precip:       day.Precip,
			PrecipType:   precipType,
			WindGust:     day.WindGust,
			WindSpeed:    day.WindSpeed,
			CloudCover:   day.CloudCover,
			Conditions:   day.Conditions,
		})
	}

	log.Debugf("Processed %d days of weather data", len(result))
	return result, nil
}
