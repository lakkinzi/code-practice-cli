package db

import (
	"log"
	"os"
	"strings"
)

const lastTaskNameFile = "last"

func SetLast(taskName string) {
	f, err := os.OpenFile(lastTaskNameFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(taskName)
	if err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func GetLast() string {
	bytes, err := os.ReadFile("./" + lastTaskNameFile)
	if err != nil {
		log.Fatal(err)
	}
	name := string(bytes)
	name = strings.Trim(strings.TrimSuffix(name, "\n"), " ")
	return name
}
