package Model

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	fs               []os.FileInfo
	CurrentDirectory string
	Pointer          int
	ShouldClose      bool
}

func Init() Model {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v", err)
		os.Exit(1)
	}

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening directory: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	fs, err := f.Readdir(-1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v", err)
		os.Exit(1)
	}

	return Model{
		fs:               fs,
		CurrentDirectory: dir,
		Pointer:          0,
		ShouldClose:      false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.ShouldClose = true
			return m, tea.Quit
		case "j", "down":
			if m.Pointer < len(m.fs)-1 {
				m.Pointer++
			}
		case "k", "up":
			if m.Pointer > 0 {
				m.Pointer--
			}
		case "right":
			if m.fs[m.Pointer].IsDir() {
				m.OpenSelectedDir()
			}
		case "left":
			m.OpenParentDir()
		}
	}
	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder

	dirStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#005cc5")).Bold(true)
	fileStyle := lipgloss.NewStyle()
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Bold(true)

	for i, v := range m.fs {
		cursor := " "
		if i == m.Pointer {
			cursor = ">"
		}

		cursor = cursorStyle.Render(cursor)

		if v.IsDir() {
			styledDir := dirStyle.Render(v.Name())
			sb.WriteString(strings.ReplaceAll(fmt.Sprintf("%s %s", cursor, styledDir), "\n", "\n  "))
			sb.WriteRune('\n')
		} else {
			sb.WriteString(fileStyle.Render(fmt.Sprintf("%s %s\n", cursor, v.Name())))
		}
	}
	return sb.String()
}
