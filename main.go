package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

var bwm BWManager

func main() {
	cmd := exec.Command("bw", "-v")
	if err := cmd.Run(); err != nil {
		fmt.Println("Could not find 'bw' command in '$PATH'. Please check if Bitwarden CLI is installed.\nGoodbye")
		os.Exit(1)
	}

	m := NewMainModel()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
