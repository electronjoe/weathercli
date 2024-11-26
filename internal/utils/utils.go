package utils

import (
	"fmt"
	"time"
)

func ParseDate(date string) (time.Time, error) {
    t, err := time.Parse("2006-01-02", date)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid date format. Please use YYYY-MM-DD format")
    }
    return t, nil
}
