package cfg

import "os"

const apiKeyFieldName = "OMDB_API_KEY"

func GetAPIKey() (string, bool) {
	return os.LookupEnv(apiKeyFieldName)
}
