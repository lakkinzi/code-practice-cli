package api

import (
	"io"
	"log"
	"net/http"
	"time"
)

var (
	codeWarsApi = "https://www.codewars.com/api/v1/"
	challanges  = "code-challenges/"
)

func Get(id string) []byte {
	url := codeWarsApi + challanges + id
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyBytes
}
