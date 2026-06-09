package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/mengstabketemaw/git-spread/models"
)

func ReadRawLog() ([]byte, error) {
	return exec.Command("git", "log", "--pretty=format:%H|%ct|%s").Output()
}

func ReadCommits() ([]models.Commit, error) {

	var commitList []models.Commit

	logs, err := exec.Command("git", "log", "--pretty=format:%H|%ct|%s").Output()

	if err != nil {
		return nil, err
	}

	lines := strings.SplitSeq(string(logs), "\n")

	for line := range lines {

		parts := strings.SplitN(line, "|", 3)

		if len(parts) != 3 {
			return nil, errors.New("invalid git log line")
		}

		timestamp, err := strconv.ParseInt(parts[1], 10, 64)

		if err != nil {
			return nil, fmt.Errorf("failed to prase timestamp: %w", err)
		}

		dateObject := time.Unix(timestamp, 0)

		commit := models.Commit{
			Hash:    parts[0],
			Message: parts[2],
			Date:    dateObject,
		}

		commitList = append(commitList, commit)
	}

	return commitList, nil
}
