package cli

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	columnKeyName = "name"
)

type modelSelect struct {
	input    textinput.Model
	selected string
	table    table.Model
}

func (m modelSelect) Init() tea.Cmd {
	return textinput.Blink
}

func createModelSelect(tasksNames []string) modelSelect {
	columns := []table.Column{
		table.NewColumn(columnKeyName, "Name", 40).WithStyle(
			lipgloss.NewStyle().
				Faint(true).
				Align(lipgloss.Left).
				Foreground(lipgloss.Color("#88f"))).
			WithFiltered(true),
	}

	rows := make([]table.Row, len(tasksNames))
	for i := range tasksNames {
		rows[i] = table.NewRow(
			table.RowData{
				columnKeyName: tasksNames[i],
			},
		)
	}

	return modelSelect{
		table: table.
			New(columns).
			WithRows(rows).
			Focused(true).
			Filtered(true).
			WithSelectedText(" ", "*"),
		input: textinput.New(),
	}
}

func (m modelSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			cmds = append(cmds, tea.Quit)

			return m, tea.Batch(cmds...)
		}

		if m.input.Focused() {
			if msg.String() == "enter" {
				m.input.Blur()
			} else {
				m.input, _ = m.input.Update(msg)
			}
			m.table = m.table.WithFilterInput(m.input)

			return m, tea.Batch(cmds...)
		}

		switch msg.String() {
		case "/":
			m.input.Focus()
		case "enter":
			m.selected = m.table.HighlightedRow().Data[columnKeyName].(string)
			return m, tea.Quit
		case "q":
			cmds = append(cmds, tea.Quit)
			return m, tea.Batch(cmds...)
		default:
			m.table, cmd = m.table.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m modelSelect) View() string {
	body := strings.Builder{}
	body.WriteString(m.input.View() + "\n")
	body.WriteString(m.table.View())
	body.WriteString("\n")

	return body.String()
}
