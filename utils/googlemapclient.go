package utils

import (
	"fmt"

	"googlemaps.github.io/maps"
)

func InitializeGoogleApi() maps.Client {
	googleclient, err := maps.NewClient(maps.WithAPIKey("AIzaSyCNYxu7FcuVSYRaRZxumdf1JnyvvTa9F38"))
	if err != nil {
		fmt.Print(err)
	}

	return *googleclient
}
