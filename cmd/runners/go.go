package runners

import (
	"fmt"
	"log"
	"os/exec"
)

func runGo(path string) {
	cmd := exec.Command("/bin/sh", "-c", "go run main.go")
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
