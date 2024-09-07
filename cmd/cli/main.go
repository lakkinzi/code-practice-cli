package cli

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

var taskID = ""

func Create() string {
	_, err := tea.NewProgram(createModelCreate()).Run()
	if err != nil {
		log.Fatal(err)
	}
	return taskID
}

func Select(tasks []string) string {
	m, err := tea.NewProgram(createModelSelect(tasks)).Run()
	if err != nil {
		log.Fatal(err)
	}

	model, ok := m.(modelSelect)
	if !ok {
		return ""
	}

	return model.selected
}
