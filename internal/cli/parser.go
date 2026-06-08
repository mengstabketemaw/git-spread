package cli

import (
	"flag"
	"fmt"
	"time"

	"github.com/mengstabketemaw/git-spread/internal/config"
)

func Parser()(config.Config, error) {
	start := flag.String("start", "", "start date")
	end := flag.String("end", "", "end date")
	dryRun := flag.Bool("dry-run", false, "dry run")

	flag.Parse()

	startDate, err := time.Parse("2006-01-02", *start)

	if err != nil {
		return config.Config{}, err
	}

	endDate, err := time.Parse("2006-01-02", *end)
	
	if err != nil {
		return config.Config{}, err
	}

	if endDate.Before(startDate) {
		return config.Config{}, fmt.Errorf("--end must be after --start")
	}

	cfg := config.Config{
		StartDate: startDate,
		EndDate: endDate,
		DryRun: *dryRun,
	}

	return cfg, nil
}
