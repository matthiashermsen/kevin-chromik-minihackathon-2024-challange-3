package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/matthiashermsen/kevin-chromik-minihackathon-2024-challange-3/cli"
)

func main() {
	err := cli.RootCommand.Execute()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("Failed execute command"))
		fmt.Println(err)
	}
}
