package formatter

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/electronjoe/weathercli/internal/api"
)

func OutputWeatherData(w io.Writer, data []api.WeatherData) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

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
	fmt.Fprintln(tw, strings.Join(headers, "\t"))

	// Write data rows
	for _, d := range data {
		fmt.Fprintf(tw, "%s\t%.1f\t%.1f\t%.1f\t%.1f\t%.1f\t%s\t%.1f\t%.1f\t%.1f\t%s\n",
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

	return tw.Flush()
}
