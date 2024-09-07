package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"practice/cmd/api"
	"practice/cmd/cli"
	"practice/cmd/config"
	"practice/cmd/db"
	"practice/cmd/editors"
	"practice/cmd/runners"
	"practice/cmd/task"
)

func getTasksPath(conf *config.Config, taskPath string) string {
	if taskPath == "" {
		return filepath.Join(conf.TasksPath, string(conf.ProgLang))
	}
	return filepath.Join(conf.TasksPath, string(conf.ProgLang), taskPath)
}

func Create(conf *config.Config) {
	id := cli.Create()
	bodyBytes := api.Get(id)
	t, err := task.Create(conf, bodyBytes)
	if err != nil {
		log.Fatal(err)
	}
	editors.Open(conf, getTasksPath(conf, t.NameCamel))
	db.SetLast(t.NameCamel)
}

func Select(conf *config.Config) {
	// tasksPath := getTasksPath(conf, "")
	files, err := os.ReadDir(conf.TasksPath)
	if err != nil {
		log.Fatal(err)
	}

	tasksNames := make([]string, len(files))
	for i := range files {
		tasksNames[i] = files[i].Name()
	}

	taskName := cli.Select(tasksNames)
	fmt.Println(taskName)
	path := filepath.Join(conf.TasksPath, taskName, string(conf.ProgLang))
	editors.Open(conf, path)
	db.SetLast(taskName)
}

func Last(conf *config.Config) {
	editors.Open(conf, getTasksPath(conf, db.GetLast()))
}

func Run(conf *config.Config) {
	runners.Run(conf, getTasksPath(conf, db.GetLast()))
}
