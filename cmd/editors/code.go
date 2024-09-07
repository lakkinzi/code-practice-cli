package editors

import (
	"fmt"
	"log"
	"os/exec"
)

func openVsCode(path string) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("code %s", path))
	cmd.Dir = path
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
