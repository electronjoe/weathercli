package formatter

import (
	"fmt"
	"io"
	"strings"

	"github.com/electronjoe/weathercli/internal/api"
)

func OutputWeatherData(w io.Writer, data []api.WeatherData) error {
	// Write header
	headers := []string{
		"Date",
		"tempmax",
		"feelslikemax",
		"tempmin",
		"feelslikemin",
		"precip",
		"preciptype",
		"windgust",
		"windspeed",
		"cloudcover",
		"conditions",
	}
	fmt.Fprintln(w, strings.Join(headers, "\t"))

	// Write data rows
	for _, d := range data {
		fmt.Fprintf(w, "%s\t%.1f\t%.1f\t%.1f\t%.1f\t%.1f\t%s\t%.1f\t%.1f\t%.1f\t%s\n",
			d.Date.Format("2006-01-02"),
			d.TempMax,
			d.FeelsLikeMax,
			d.TempMin,
			d.FeelsLikeMin,
			d.Precip,
			d.PrecipType,
			d.WindGust,
			d.WindSpeed,
			d.CloudCover,
			d.Conditions,
		)
	}

	return nil
}
