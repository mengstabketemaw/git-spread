package main

import (
	"fmt"

	"github.com/mengstabketemaw/git-spread/internal/cli"
	"github.com/mengstabketemaw/git-spread/internal/git"
	"github.com/mengstabketemaw/git-spread/internal/rewrite"
	"github.com/mengstabketemaw/git-spread/internal/scheduler"
)

func main() {

	cfg, err := cli.Parser()

	if err != nil {
		x := 1
		fmt.Println(x)
		return
	}

	commits, err := git.ReadCommits()

	if err != nil {
		fmt.Println(err)
	}

	scheduled, err := scheduler.Schedule(commits, cfg)

	if err != nil {
		fmt.Print(err)
	}

	if cfg.DryRun {
		fmt.Printf("%-10s %-15s %-15s\n", "Hash", "Old", "New")
		for i := range scheduled {
			fmt.Printf("%-10s %-15s %-15s\n", scheduled[i].Hash, commits[i].Date, scheduled[i].Date)
		}
	}
	fmt.Println(commits, err)

	rewrite.Test()

}
