package main

import (
	"fmt"
	"time"
)

func ToMinut(timee string) (string, error) {
	duration, err := time.ParseDuration(timee)
	if err != nil {
		return "", fmt.Errorf("failed to make conversion %v", err)
	}
	return fmt.Sprintf("%s = %v min", timee, duration.Minutes()), nil
}
