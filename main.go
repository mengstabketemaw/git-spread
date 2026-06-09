package main

import (
	"fmt"

	"github.com/mengstabketemaw/git-spread/internal/cli"
	"github.com/mengstabketemaw/git-spread/internal/git"
)

func main(){

	cfg, err := cli.Parser()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(cfg)

	git.ReadCommits()
	
}
