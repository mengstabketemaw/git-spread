package main

import (
	"fmt"
	"github.com/mengstabketemaw/git-spread/internal/cli"
)

func main(){

	cfg, err := cli.Parser()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(cfg)
}
