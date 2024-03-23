package choose

import (
	"fmt"
	"os"
	"strings"

	"github.com/KevinYouu/fastGit/functions/colors"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"fix", "feat", "docs", "style", "refactor", "test", "chore", "revert"}

type model struct {
	cursor   int
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter", " ":
			m.quitting = true
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("What kind of Bubble Tea would you like to order?\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString(colors.RenderColor("blue", "◉ "+choices[i]))
		} else {
			s.WriteString(colors.RenderColor("white", "○ "+choices[i]))
		}
		s.WriteString("\n")
	}
	s.WriteString("\x1b[0m") // reset color
	s.WriteString("(press q to quit)\n")
	if m.quitting {
		return "" // clear the screen
	}
	return s.String()
}

func Choose() string {
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		// fmt.Printf("\n---\nYou chose %s!\n", m.choice)
		return m.choice
	}

	return ""
}
