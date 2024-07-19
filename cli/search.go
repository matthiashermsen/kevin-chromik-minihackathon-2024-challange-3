package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"github.com/matthiashermsen/kevin-chromik-minihackathon-2024-challange-3/cfg"
)

var yearOfReleaseFlag string

func init() {
	SearchCommand.Flags().StringVarP(&yearOfReleaseFlag, "yearOfRelease", "y", "", "year of release")

	RootCommand.AddCommand(SearchCommand)
}

var SearchCommand = &cobra.Command{
	Use:   "search <title> [--yearOfRelease | -y <year>]",
	Short: "Search by title and year of release ( optional )",
	Long:  "Search movie or series by title and year of release ( optional ).",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		apiKey, hasAPIKey := cfg.GetAPIKey()

		if !hasAPIKey {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("API key is missing. Please add 'OMDB_API_KEY={your_api_key}' to the environment variables"))

			return
		}

		searchTitle := args[0]
		queryParams := url.Values{
			"apikey": {apiKey},
			"s":      {searchTitle},
		}

		if yearOfReleaseFlag != "" {
			queryParams.Add("y", yearOfReleaseFlag)
		}

		encodedQueryParams := queryParams.Encode()
		apiURL := fmt.Sprintf("https://www.omdbapi.com/?%s", encodedQueryParams)

		apiResponse, err := http.Get(apiURL)
		if err != nil {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("Failed to fetch data"))
			fmt.Println(err)

			return
		}
		defer apiResponse.Body.Close()

		if apiResponse.StatusCode != http.StatusOK {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("Request failed"))

			return
		}

		var searchAPIResponse SearchAPIResponse
		err = json.NewDecoder(apiResponse.Body).Decode(&searchAPIResponse)
		if err != nil {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("Failed to parse response body"))
			fmt.Println(err)

			return
		}

		if searchAPIResponse.IsResponseSuccessful != "True" {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("Request failed"))
			fmt.Println(searchAPIResponse.ErrorMessage)

			return
		}

		if searchAPIResponse.AmoutOfResults == "1" {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render("Success! Found 1 entry"))
		} else {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render(fmt.Sprintf("Success! Found %s entries", searchAPIResponse.AmoutOfResults)))
		}

		for _, result := range searchAPIResponse.Results {
			boxStyle := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#874BFD")).
				Background(lipgloss.Color("#000000")).
				Padding(1, 0).
				BorderTop(true).
				BorderLeft(true).
				BorderRight(true).
				BorderBottom(true).
				Margin(2, 0).
				Padding(1, 2)

			titleStyle := lipgloss.NewStyle().Bold(true)
			typeAndYearStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#696969"))
			imdbReferenceStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#f6e100")).Margin(1, 0)
			posterReferenceStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#898989"))
			moreInformationSuggestionStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#96ffec")).MarginTop(1)

			fmt.Println(boxStyle.Render(fmt.Sprintf(
				"%s\n%s\n%s\n%s\n%s",
				titleStyle.Render(result.Title),
				typeAndYearStyle.Render(fmt.Sprintf("%s | %s", result.Type, result.Year)),
				imdbReferenceStyle.Render(fmt.Sprintf("Reference: https://www.imdb.com/title/%s", result.IMDBID)),
				posterReferenceStyle.Render(fmt.Sprintf("Poster: %s", result.PosterURL)),
				moreInformationSuggestionStyle.Render(fmt.Sprintf("Get more information by running 'omdb get %s'", result.IMDBID)),
			)))
		}
	},
}
