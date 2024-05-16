package main

import (
	"fmt"
	"os"
	"zfm/Model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := Model.Init()

	tm := tea.NewProgram(&p, tea.WithAltScreen())

	m, err := tm.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v", err)
		os.Exit(1)
	}

	if m.(Model.Model).ShouldClose {
		fmt.Println("Goodbye!")
	}
}
