package editors

import (
	"fmt"
	"log"
	"os/exec"
)

func openNvim(path string) {
	p1 := "--working-directory=" + path
	// cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("terminator %s -x nvim .", p1))
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("alacritty %s -e nvim .", p1))
	cmd.Dir = path
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
