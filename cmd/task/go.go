package task

import (
	"os"
	"path/filepath"
	"practice/cmd/config"
	"practice/cmd/files"
)

func createGo(conf *config.Config, task *Task) error {
	err := createGoFiles(conf, task)
	if err != nil {
		return err
	}
	return nil
}

func createGoFiles(conf *config.Config, task *Task) error {
	path := filepath.Join(conf.TasksPath, string(config.Go), task.NameCamel)

	err := os.Mkdir(path, 0o777)
	if err != nil {
		return nil
	}

	taskFilePath := path + "/main.go"
	templatePath := "cmd/task/templates/main.tmpl"
	err = files.CreateFileFromTemplate(templatePath, taskFilePath, task)
	if err != nil {
		return nil
	}

	taskFilePath = path + "/description.md"
	templatePath = "cmd/task/templates/description.tmpl"
	err = files.CreateFileFromTemplate(templatePath, taskFilePath, task)
	if err != nil {
		return nil
	}

	return nil
}
