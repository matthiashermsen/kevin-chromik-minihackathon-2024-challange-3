package cli

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "omdb",
	Short: "omdb is a tool to search for movies and series",
	Long:  "omdb is a tool to search for movies and series.",
}
