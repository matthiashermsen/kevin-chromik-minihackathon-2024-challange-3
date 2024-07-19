package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/charmbracelet/lipgloss"
	"github.com/matthiashermsen/kevin-chromik-minihackathon-2024-challange-3/cfg"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(GetByIMDBIDCommand)
}

var GetByIMDBIDCommand = &cobra.Command{
	Use:   "get <imdbID>",
	Short: "Get by IMDB ID",
	Long:  "Get movie or series by IMDB ID.",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		apiKey, hasAPIKey := cfg.GetAPIKey()

		if !hasAPIKey {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("API key is missing. Please add 'OMDB_API_KEY={your_api_key}' to the environment variables"))

			return
		}

		imdbID := args[0]
		queryParams := url.Values{
			"apikey": {apiKey},
			"i":      {imdbID},
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

		var getAPIResponse GetAPIResponse
		err = json.NewDecoder(apiResponse.Body).Decode(&getAPIResponse)
		if err != nil {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("Failed to parse response body"))
			fmt.Println(err)

			return
		}

		if getAPIResponse.IsResponseSuccessful != "True" {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("Request failed"))
			fmt.Println(getAPIResponse.ErrorMessage)

			return
		}

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
		typeAndYearWithReleaseStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#696969"))
		runtimeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#696969"))
		ratingsStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#96ffec")).MarginTop(1)
		genreStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#696969")).MarginTop(1)
		plotStyle := lipgloss.NewStyle().MarginTop(1)
		directorAndWriterStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#696969")).MarginTop(1)
		actorsStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#696969"))
		countryStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#696969")).MarginTop(1)
		languageStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#696969"))
		posterReferenceStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#898989"))

		fmt.Println(boxStyle.Render(fmt.Sprintf(
			"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
			titleStyle.Render(getAPIResponse.Title),
			typeAndYearWithReleaseStyle.Render(fmt.Sprintf("%s | %s ( released %s )", getAPIResponse.Type, getAPIResponse.Year, getAPIResponse.Released)),
			runtimeStyle.Render(fmt.Sprintf("Runtime: %s", getAPIResponse.Runtime)),
			ratingsStyle.Render(fmt.Sprintf("Metascore: %s | IMDB: %s ( %s votes )", getAPIResponse.Metascore, getAPIResponse.IMDBRating, getAPIResponse.IMDBVotes)),
			genreStyle.Render(fmt.Sprintf("Genre: %s", getAPIResponse.Genre)),
			plotStyle.Render(getAPIResponse.Plot),
			directorAndWriterStyle.Render(fmt.Sprintf("Director: %s | Writer: %s", getAPIResponse.Director, getAPIResponse.Writer)),
			actorsStyle.Render(fmt.Sprintf("Actors: %s", getAPIResponse.Actors)),
			countryStyle.Render(fmt.Sprintf("Country: %s", getAPIResponse.Country)),
			languageStyle.Render(fmt.Sprintf("Language: %s", getAPIResponse.Language)),
			posterReferenceStyle.Render(fmt.Sprintf("Poster: %s", getAPIResponse.PosterURL)),
		)))
	},
}
