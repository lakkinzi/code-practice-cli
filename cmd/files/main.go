package files

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func CreateFile(name string) (*os.File, error) {
	return os.Create(name + ".json")
}

func WriteToFile(name string, text string) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text)
	return err
}

func GetBody(r *http.Response) (string, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString, err
}

func CreateFileFromTemplate(templatePath string, filePath string, data interface{}) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	taskFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	err = tmpl.Execute(taskFile, data)
	if err != nil {
		return err
	}
	return nil
}
