package config

import "time"

type Config struct {
	StartDate time.Time
	EndDate time.Time

	DryRun bool
}
