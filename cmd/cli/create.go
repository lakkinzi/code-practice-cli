package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

type modelCreate struct {
	focused int
	err     error
	inputs  []textinput.Model
}

func (m modelCreate) Init() tea.Cmd {
	return textinput.Blink
}

func createModelCreate() modelCreate {
	inputs := make([]textinput.Model, 1)
	inputs[0] = textinput.New()
	inputs[0].Placeholder = "ID"
	inputs[0].Focus()

	return modelCreate{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m modelCreate) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			taskID = m.inputs[0].Value()
			return m, tea.Quit
		}
	}
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

var inputStyle = lipgloss.NewStyle()

func (m modelCreate) View() string {
	s := "FindTask: \n"
	s += fmt.Sprintf(`
%s
%s
`,
		inputStyle.Width(30).Render("Task"),
		m.inputs[0].View(),
	)

	s += "\nPress q to quit.\n"

	return s
}
