package rewrite

import (
	"fmt"
	"os"
	"os/exec"
)

func Test() {
	cmd := exec.Command("env")
	cmd.Env = append(
		os.Environ(),
		"APPNAME=git-spread",
		)

	output, _ := cmd.Output()

	fmt.Println(string(output))
}
