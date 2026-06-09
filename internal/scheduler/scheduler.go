package scheduler

import (
	"time"

	"github.com/mengstabketemaw/git-spread/internal/config"
	"github.com/mengstabketemaw/git-spread/models"
)

func Schedule(commits []models.Commit, cfg config.Config) ([]models.Commit, error) {
	newCommits := make([]models.Commit, len(commits))

	copy(newCommits, commits)

	if len(commits) < 2 {
		return newCommits, nil
	}

	duration := cfg.EndDate.Sub(cfg.StartDate)
	step := duration / time.Duration(len(commits)-1)

	for i := range newCommits {
		newCommits[i].Date = cfg.StartDate.Add(step * time.Duration(i))
	}

	return newCommits, nil
}
