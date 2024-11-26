package cmd

import (
	"fmt"
	"os"

	"github.com/electronjoe/weathercli/internal/api"
	"github.com/electronjoe/weathercli/internal/config"
	"github.com/electronjoe/weathercli/internal/formatter"
	"github.com/electronjoe/weathercli/internal/utils"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.WarnLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

var rootCmd = &cobra.Command{
	Use:   "weathercli [city] [start-date]",
	Short: "A CLI tool to fetch and display weather data",
	Long: `weathercli is a command-line tool that fetches historical weather data
from VisualCrossing's Weather API and displays it in a tabular format.
Example: weathercli "London, UK" 2024-03-15`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get city and validate it's not empty
		city := args[0]
		if city == "" {
			return fmt.Errorf("city parameter cannot be empty")
		}

		// Validate and parse start date
		startDate, err := utils.ParseDate(args[1])
		if err != nil {
			return fmt.Errorf("invalid date format: %v", err)
		}

		// Get API configuration
		cfg, err := config.NewConfig()
		if err != nil {
			return err
		}

		// Create API client
		client := api.NewClient(cfg.APIKey)

		// Fetch weather data
		data, err := client.FetchWeatherData(city, startDate)
		if err != nil {
			return err
		}

		// Format and output data
		return formatter.OutputWeatherData(os.Stdout, data)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
