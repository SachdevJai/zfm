package main

import (
	"fmt"
	"os"
	"zfm/internal"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := internal.Init()

	tm := tea.NewProgram(&p, tea.WithAltScreen())

	_, err := tm.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v", err)
		os.Exit(1)
	}

}
