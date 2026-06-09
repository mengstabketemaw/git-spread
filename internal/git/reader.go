package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mengstabketemaw/git-spread/models"
)

func ReadRawLog() ([]byte, error) {
	return exec.Command("git", "log", "--pretty=format:%H|%ct|%s").Output()
}

func ReadCommits() ([]models.Commit, error){
	logs, err := exec.Command("git", "log", "--pretty=format:%H|%ct|%s").Output()

	if err != nil {
		return []models.Commit{}, err
	}

	parts := strings.Split(string(logs), "|")

	fmt.Print(parts)

	return nil, nil
}
