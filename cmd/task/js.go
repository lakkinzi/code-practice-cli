package task

import (
	"os"
	"path/filepath"
	"practice/cmd/config"
	"practice/cmd/files"
)

func createJS(conf *config.Config, task *Task) error {
	err := createJsFiles(conf, task)
	if err != nil {
		return err
	}
	return nil
}

func createJsFiles(conf *config.Config, task *Task) error {
	path := filepath.Join(conf.TasksPath, string(config.Javascript), task.NameCamel)

	err := os.Mkdir(path, 0o777)
	if err != nil {
		return nil
	}

	taskFilePath := path + "/index.js"
	templatePath := "cmd/task/templates/js.tmpl"
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
