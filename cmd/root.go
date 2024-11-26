package cmd

import (
	"fmt"
	"os"

	"github.com/electronjoe/weathercli/internal/api"
	"github.com/electronjoe/weathercli/internal/config"
	"github.com/electronjoe/weathercli/internal/formatter"
	"github.com/electronjoe/weathercli/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "weathercli [start-date]",
	Short: "A CLI tool to fetch and display weather data",
	Long: `weathercli is a command-line tool that fetches historical weather data
from VisualCrossing's Weather API and displays it in a tabular format.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate and parse start date
		startDate, err := utils.ParseDate(args[0])
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
		data, err := client.FetchWeatherData(startDate)
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
