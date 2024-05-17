package internal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func renderFilePanel(m *Model) string {

	var sb strings.Builder

	dirStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#005cc5")).Bold(true)
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Bold(true)

	boxWidth := int(float32(m.termWidth) * 0.99)
	boxHeight := m.termHeight - 5

	offset := 0

	if m.pointer > boxHeight-1 {
		offset = m.pointer - boxHeight + 1
	}

	for i := 0; i < boxHeight; i++ {

		if (i + offset) >= len(m.fs) {
			break
		}

		v := m.fs[i+offset]
		cursor := " "
		if i+offset == m.pointer {
			cursor = ">"
		}

		cursor = cursorStyle.Render(cursor)

		var styledName string
		if v.IsDir() {
			styledName = dirStyle.Render(v.Name())
		} else {
			styledName = v.Name()
		}

		if i+offset == m.pointer {
			styledName = cursorStyle.Render(styledName)
		}

		sb.WriteString(fmt.Sprintf("%s %s\n", cursor, styledName))

	}

	boxStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#00ff00")).
		Width(boxWidth).
		Height(boxHeight)

	panel := boxStyle.Render(sb.String())

	return panel
}

func renderCommandLine(m *Model) string {

	boxWidth := int(float32(m.termWidth) * 0.99)
	boxHeight := 1

	boxStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#00ff00")).
		Width(boxWidth).
		Height(boxHeight)

	cmdLine := fmt.Sprintf(": %s", m.commandBuffer)

	if !m.commandMode {
		cmdLine = lipgloss.NewStyle().Foreground(lipgloss.Color("#005cc5")).Render(cmdLine)
	} else {
		cmdLine = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(cmdLine)
	}

	panel := boxStyle.Render(cmdLine)

	return panel
}
