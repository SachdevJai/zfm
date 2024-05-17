package internal

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type Model struct {
	fs               []os.FileInfo
	currentDirectory string
	pointer          int
	shouldClose      bool
	termWidth        int
	termHeight       int
	commandBuffer    string
	commandMode      bool
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

	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v", err)
		os.Exit(1)
	}

	return Model{
		fs:               fs,
		currentDirectory: dir,
		pointer:          0,
		shouldClose:      false,
		termWidth:        width,
		termHeight:       height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.commandMode {
			switch msg.String() {
			case "enter":
				executeCommand(m.commandBuffer)
				m.commandBuffer = ""
				m.commandMode = false
			case "ctrl + c":
				m.commandBuffer = ""
				m.commandMode = false
			case "backspace":
				m.commandBuffer = m.commandBuffer[:len(m.commandBuffer)-1]
			default:
				m.commandBuffer += msg.String()
			}

		} else {
			switch msg.String() {
			case "ctrl+c", "q":
				m.shouldClose = true
				return m, tea.Quit
			case "j", "down":
				if m.pointer < len(m.fs)-1 {
					m.pointer++
				}
			case "k", "up":
				if m.pointer > 0 {
					m.pointer--
				}
			case "right":
				if m.fs[m.pointer].IsDir() {
					m.OpenSelectedDir()
				}
			case "left":
				m.OpenParentDir()
			case ":":
				m.commandBuffer = ""
				m.commandMode = true
			}
		}
	}
	return m, nil
}

func (m Model) View() string {

	filePanel := renderFilePanel(&m)
	commandLine := renderCommandLine(&m)

	mainPanel := lipgloss.JoinVertical(0, filePanel, commandLine)

	return mainPanel
}
