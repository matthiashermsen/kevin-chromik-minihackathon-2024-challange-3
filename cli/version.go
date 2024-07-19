package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"github.com/matthiashermsen/kevin-chromik-minihackathon-2024-challange-3/app"
)

func init() {
	RootCommand.AddCommand(VersionCommand)
}

var VersionCommand = &cobra.Command{
	Use:   "version",
	Short: "Prints the current app version",
	Long:  "Prints the current app version.",
	Run: func(_ *cobra.Command, _ []string) {
		if app.Version == "" {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("Version is not available"))

			return
		}

		fmt.Println(app.Version)
	},
}
