package Model

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
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
		}
	}
	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder

	_, height, err := term.GetSize(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v", err)
		os.Exit(1)
	}

	for i := 0; i < height; i++ {
		sb.WriteString("\n")
	}

	if len(m.fs) == 0 {
		return "No files"
	}

	for i, f := range m.fs {
		cursor := " "
		if i == m.Pointer {
			cursor = ">"
		}
		sb.WriteString(cursor)
		sb.WriteString(f.Name())
		sb.WriteString("\n")
	}

	return sb.String()
}
