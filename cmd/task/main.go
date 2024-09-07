package task

import (
	"encoding/json"
	"errors"
	"practice/cmd/config"
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	NameCamel   string
	ProgLang    config.ProgLang
}

var constructors = map[config.ProgLang]func(*config.Config, *Task) error{
	config.Go:         createGo,
	config.Javascript: createJS,
}

func Create(conf *config.Config, bytes []byte) (*Task, error) {
	task, err := createTask(bytes)
	if err != nil {
		return nil, err
	}

	constructor, ok := constructors[conf.ProgLang]
	if !ok {
		return nil, errors.New("not language")
	}

	err = constructor(conf, task)
	if err != nil {
		return nil, err
	}

	return task, err
}

func createTask(bytes []byte) (*Task, error) {
	task := &Task{}
	err := json.Unmarshal(bytes, task)
	task.NameCamel = ToLowerCamel(task.Name)
	if err != nil {
		return nil, err
	}
	return task, err
}
