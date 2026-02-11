// Package main provides a simple example of how to use the generated SDK.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

func main() {
	apiKey := os.Getenv("RETOOL_ACCESS_TOKEN")
	host := os.Getenv("RETOOL_HOST")
	scheme := os.Getenv("RETOOL_SCHEME")

	if apiKey == "" {
		fmt.Println("Please set the RETOOL_ACCESS_TOKEN environment variable")
		return
	}

	if host == "" {
		host = "localhost:3000"
	}

	if scheme == "" {
		scheme = "http"
	}

	config := api.NewConfiguration()
	config.Host = host
	config.Scheme = scheme
	config.Servers = api.ServerConfigurations{api.ServerConfiguration{
		URL: "/api/v2",
	},
	}
	client := api.NewAPIClient(config)
	auth := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)
	folders, httpResponse, err := client.FoldersAPI.FoldersGet(auth).Execute()
	if err != nil {
		log.Fatal("Error!", err, httpResponse)
	}
	fmt.Println("Folders:")
	for _, folder := range folders.Data {
		fmt.Printf("%s (%s)\n", folder.Name, folder.FolderType)
	}
}
